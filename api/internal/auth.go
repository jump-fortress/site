package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rotisserie/eris"
	"github.com/spiritov/jump/api/db/queries"
	db "github.com/spiritov/jump/api/db/responses"
	"github.com/spiritov/jump/api/env"
	"github.com/spiritov/jump/api/slog"
	"github.com/yohcop/openid-go"
)

const (
	SessionCookieName = "sessionid"

	SessionIssuer       = "jump"
	SessionAudience     = "jump"
	SessionDuration     = time.Hour * 24 * 7
	SessionJitter       = time.Minute
	PrincipalContextKey = "principal"

	SteamOidcIssuer      = "https://steamcommunity.com/openid/"
	SteamOidRedirectPath = "/session/steam/callback"
)

var (
	SessionTokenSecret  []byte
	SessionCookieSecure = false
	OidRealm            string
	OidRealmURL         *url.URL
	SteamApiKey         string

	discoveryCache *NoOpDiscoveryCache
)

type Principal struct {
	SteamID uint64
	TokenID uuid.UUID
	Claims  *jwt.RegisteredClaims
}

func GetPrincipal(ctx context.Context) (result *Principal, ok bool) {
	result, ok = ctx.Value(PrincipalContextKey).(*Principal)
	ok = ok && result != nil
	return
}

func HasPrincipal(ctx context.Context) bool {
	result, ok := ctx.Value(PrincipalContextKey).(*Principal)
	return ok && result != nil
}

type DiscoverInput struct {
	URL url.URL
}

func (d *DiscoverInput) Resolve(ctx huma.Context) []error {
	d.URL = ctx.URL()
	return nil
}

type DiscoverOutput struct {
	Status int
	Url    string `header:"Location"`
}

func handleSteamDiscover(ctx context.Context, input *DiscoverInput) (*DiscoverOutput, error) {
	callbackURL := OidRealmURL.JoinPath(SteamOidRedirectPath)
	redirectUrl, err := openid.RedirectURL(SteamOidcIssuer, callbackURL.String(), OidRealm)

	if err != nil {
		log.Printf("[error] couldn't create openid redirect: %v", err)
		return nil, err
	}

	return &DiscoverOutput{
		Status: http.StatusTemporaryRedirect,
		Url:    redirectUrl,
	}, nil
}

type CallbackInput struct {
	URL url.URL
}

func (d *CallbackInput) Resolve(ctx huma.Context) []error {
	d.URL = ctx.URL()
	return nil
}

type CallbackOutput struct {
	Status    int
	Url       string      `header:"Location"`
	SetCookie http.Cookie `header:"Set-Cookie"`
}

func handleSteamCallback(ctx context.Context, input *CallbackInput) (*CallbackOutput, error) {
	// our openid library verifies that the original request came from our authority, but it needs us to
	// provide a URL to verify that the incoming callback request has the authority we expect. Here we're
	// just replacing the `https://blahblah.com` part of the URL with our OidRealm
	inputURL := input.URL
	fullURL := OidRealmURL.JoinPath(inputURL.Path)
	fullURL.RawQuery = inputURL.RawQuery

	// verify the openid callback. The discovery cache caches some response information to make verification
	// faster if the callback is hit with the same user again. The nonce store ensures that a callback request
	// is never processed by our servers more than once.
	id, err := openid.Verify(fullURL.String(), discoveryCache, db.NewNonceStore(ctx, db.Queries))
	if err != nil {
		slog.Logger.Debug("Error verifying openid callback", "error", err, "uri", fullURL)
		return nil, eris.Wrap(err, "Error verifying openid callback")
	}

	slog.Logger.Debug("Verified openid callback for steam user", "user_id", id)

	// the openid id is in the format `https?://steamcommunity.com/openid/id/[0-9]+`. We only care about the
	// last part, which is the user's Steam ID 64.
	var steamID string
	if strings.HasPrefix(id, "https") {
		_, err = fmt.Sscanf(id, "https://steamcommunity.com/openid/id/%s", &steamID)
	} else {
		_, err = fmt.Sscanf(id, "http://steamcommunity.com/openid/id/%s", &steamID)
	}

	if err != nil {
		slog.Logger.Error("Verified openid callback but couldn't parse Steam ID 64 from ID.", "id", id, "error", err)
		return nil, eris.Wrap(err, "Error parsing Steam ID 64 from ID")
	}

	// we need the steam ID as a string, but we want to ensure that it is a valid uint64 first
	if _, parseErr := strconv.ParseUint(steamID, 10, 64); parseErr != nil {
		slog.Logger.Error("Verified openid callback but couldn't parse Steam ID 64 from ID.", "id", id, "error", parseErr)
		return nil, eris.Wrap(parseErr, "Error parsing Steam ID 64")
	}

	// if the user doesn't already exist in the database, we need to ensure they exist
	_, dbErr := db.Queries.InsertPlayer(ctx, steamID)
	if dbErr != nil {
		return nil, eris.Wrap(dbErr, "Error creating new user")
	}

	// AddSession will create a new session UUIDv7 token entry and link it to the user.
	session, dbErr := db.Queries.AddSession(ctx, queries.AddSessionParams{
		TokenID:   uuid.Must(uuid.NewV7()).String(),
		SteamId64: steamID,
	})
	if dbErr != nil {
		return nil, eris.Wrap(dbErr, "Error creating new session")
	}

	// There are a handful of "claims" we need to specify in the JWT. The subject and ID are the most
	// important, since they specify the user's authenticated steam ID and the token's UUID.
	expiresAt := session.CreatedAt.Add(SessionDuration)
	claims := jwt.RegisteredClaims{
		Issuer:    SessionIssuer,
		Subject:   steamID,
		Audience:  []string{SessionAudience},
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		NotBefore: jwt.NewNumericDate(session.CreatedAt.Add(-SessionJitter)),
		IssuedAt:  jwt.NewNumericDate(session.CreatedAt),
		ID:        session.TokenID,
	}

	// this is creating and signing the actual JWT with the claims & secret we provide.
	signedJwt, jwtErr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SessionTokenSecret)
	if jwtErr != nil {
		return nil, eris.Wrap(jwtErr, "Error creating new session")
	}

	// and finally, setting the session cookie with our session JWT!
	return &CallbackOutput{
		Status: http.StatusTemporaryRedirect,
		Url:    "http://localhost:5173",
		SetCookie: http.Cookie{
			Name:     SessionCookieName,
			Path:     "/",
			Value:    signedJwt,
			MaxAge:   int(expiresAt.Sub(time.Now().UTC()).Seconds()),
			Expires:  expiresAt,
			Secure:   SessionCookieSecure,
			SameSite: http.SameSiteStrictMode,
		},
	}, nil
}

type SignOutOutput struct {
	SetCookie http.Cookie `header:"Set-Cookie"`
}

func handleSteamSignOut(ctx context.Context, _ *struct{}) (*SignOutOutput, error) {
	// if we don't have a principal, that means the user is not signed in or their session has expired.
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	// but, if we do have a session, we should forcefully invalidate the session to ensure the user's token
	// can't be re-used. TokenIDs are UUIDv7, which have a time-based monotonic counter as part of the ID...
	// as a result, it's virtually impossible for the same token to be generated twice.
	err := db.Queries.DisallowToken(ctx, principal.TokenID.String())
	if err != nil {
		return nil, eris.Wrap(err, "Error signing out session")
	}

	// then we expire their session cookie.
	return &SignOutOutput{
		SetCookie: http.Cookie{
			Name:     SessionCookieName,
			Path:     "/",
			Value:    "",
			MaxAge:   0,
			Expires:  time.Now(),
			Secure:   SessionCookieSecure,
			SameSite: http.SameSiteStrictMode,
		},
	}, nil
}

func registerAuth(sessionApi *huma.Group, internalApi *huma.Group) {
	OidRealm = env.GetString("JUMP_OID_REALM")
	oidRealmURL, err := url.Parse(OidRealm)
	if err != nil {
		log.Fatalf("[fatal] error parsing JUMP_OID_REALM: %v", err)
	}

	OidRealmURL = oidRealmURL
	SessionTokenSecret = []byte(env.GetString("JUMP_SESSION_TOKEN_SECRET"))
	SessionCookieSecure = env.GetBool("JUMP_SESSION_COOKIE_SECURE")
	SteamApiKey = env.GetString("JUMP_STEAM_API_KEY")

	// the OpenID flow works like this:
	// - The user is redirected to `internal/session/steam/discover`
	// - /internal/session/steam/discover does some magic & redirects the user to the Steam OpenID auth flow
	// - Once the user logs in, steam redirects the user to `internal/session/steam/callback`,
	//   with some information about the user's auth session
	// - `/internal/session/steam/callback` creates a new session token for the user
	// - the user is redirected back to home with their session cookies set
	huma.Get(sessionApi, "/steam/discover", handleSteamDiscover)
	huma.Get(sessionApi, "/steam/callback", handleSteamCallback)

	var sessionCookieSecurityMap = []map[string][]string{{"Steam": {}}}
	var requireUserSessionMiddlewares = huma.Middlewares{UserAuthHandler, CreateRequireUserAuthHandler(internalApi)}

	huma.Register(sessionApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/steam/profile",
		OperationID: "steam-profile",
		Summary:     "Steam profile",
		Description: "Get the authenticated user's steam profile info",
		Errors:      []int{http.StatusUnauthorized},

		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, handleSteamProfile)

	huma.Register(sessionApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/sign-out",
		OperationID: "sign-out",
		Summary:     "Sign out",
		Description: "Sign out & clear session",
		Errors:      []int{http.StatusUnauthorized},

		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, handleSteamSignOut)
}
