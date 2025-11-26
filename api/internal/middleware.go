package internal

import (
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spiritov/jump/api/db/responses"
)

func UserAuthHandler(ctx huma.Context, next func(huma.Context)) {
	// the session cookie is a JWT. The JWT contains claims like `sub`, `jti`, `exp`
	// these claims are strings which store information about the session
	sessionCookie, cookieErr := huma.ReadCookie(ctx, SessionCookieName)
	if cookieErr != nil {
		// if at any point we fail authenticating the user, we just skip
		// to the next HTTP handler without adding any new session info
		// to the context.
		next(ctx)
		return
	}

	// we parse the JWT and verify that its been signed with SessionTokenSecret
	token, parseErr := jwt.ParseWithClaims(
		sessionCookie.Value,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return SessionTokenSecret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer(SessionIssuer),
		jwt.WithAudience(SessionAudience),
	)
	if parseErr != nil {
		next(ctx)
		return
	}

	// then we grab the subject, which should always be the user's SteamID 64
	subjectString, subjectErr := token.Claims.GetSubject()
	if subjectErr != nil {
		next(ctx)
		return
	}

	// so it should always succeed parsing to uint64
	steamId, intErr := strconv.ParseUint(subjectString, 10, 64)
	if intErr != nil {
		next(ctx)
		return
	}

	// the jti will always be a UUID. Token IDs can be used to forcefully
	// invalidate a token. This is useful when a user's account has been
	// compromised, and we need to invalidate all of their JWTs. It's also
	// important to invalidate the user's token when they sign out.

	tokenId, tokenIdErr := uuid.FromString(token.Claims.(*jwt.RegisteredClaims).ID)
	if tokenIdErr != nil {
		next(ctx)
		return
	}

	// if the user's token has been marked as disallowed, then the authentication
	// should always fail
	isDisallowed, err := responses.Queries.GetDisallowToken(ctx.Context(), tokenId.String())
	if isDisallowed == 1 || err != nil {
		next(ctx)
		return
	}

	// The huma context can't be mutated directly. Instead, we have to wrap
	// the existing context with a new context whenever we want to add new
	// key-value pairs. In this case, we want to add the "principal".
	// The Principal is the user that was authenticated during the request,
	// it can contain any information we want.
	ctx = huma.WithValue(ctx, PrincipalContextKey, &Principal{
		SteamID: steamId,
		TokenID: tokenId,
		Claims:  token.Claims.(*jwt.RegisteredClaims),
	})
	next(ctx)
}

func CreateRequireUserAuthHandler(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	// this is a function which returns a function.
	// Think of this like a constructor for a class called "RequireUserAuthHandler", which takes one
	// parameter: a huma.API. We need that huma.API in order to call huma.WriteErr whenever the user is not
	// authenticated in the request.
	//
	// also, this is called function currying! We're returning a new function that has a closure on `api` (:
	return func(ctx huma.Context, next func(huma.Context)) {
		if !HasPrincipal(ctx.Context()) {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "")
			return
		}

		next(ctx)
	}
}

func CreateRequireUserModeratorHandler(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		principal, ok := GetPrincipal(ctx.Context())
		if !ok {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "")
			return
		}
		steamID64_string := strconv.FormatUint(principal.SteamID, 10)

		player, err := responses.Queries.SelectPlayerFromSteamID64(ctx.Context(), steamID64_string)
		if err != nil {
			_ = huma.WriteErr(api, ctx, http.StatusInternalServerError, "")
			return
		}

		if player.Role != "Mod" && player.Role != "Admin" {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "")
			return
		}

		next(ctx)
	}
}

func CreateRequireUserAdminHandler(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		principal, ok := GetPrincipal(ctx.Context())
		if !ok {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "")
			return
		}
		steamID64_string := strconv.FormatUint(principal.SteamID, 10)

		player, err := responses.Queries.SelectPlayerFromSteamID64(ctx.Context(), steamID64_string)
		if err != nil {
			_ = huma.WriteErr(api, ctx, http.StatusInternalServerError, "")
			return
		}

		if player.Role != "Admin" {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "")
			return
		}

		next(ctx)
	}
}
