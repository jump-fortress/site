package internal

import (
	"context"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

var (
	api                             huma.API
	sessionCookieSecurityMap        = []map[string][]string{{"Steam": {}}}
	requireUserSessionMiddlewares   huma.Middlewares
	requireUserModeratorMiddlewares huma.Middlewares
	requireUserAdminMiddlewares     huma.Middlewares
)

func setupRouter() *chi.Mux {
	router := chi.NewMux()

	// todo: use strict `AllowedOrigins`
	// todo: use CSRF middleware (aa)
	// todo: rate limit
	router.Use(cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"}, // default value
	}).Handler)

	return router
}

func setupHumaConfig() huma.Config {
	config := huma.DefaultConfig("Jump Fortress API", "1.0.0")

	// steam security scheme, a JWT with user's OpenID information
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"Steam": {
			Type:        "apiKey",
			In:          "cookie",
			Description: "a session cookie stores the user's session token.",
			Name:        SessionCookieName,
		},
	}
	return config
}

func ServeAPI(address string) {
	router := setupRouter()
	config := setupHumaConfig()
	api = humachi.New(router, config)

	registerRoutes()

	// last
	log.Print("serving api on " + address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("[fatal] failed to serve api: %v", err)
	}
}

func registerRoutes() {
	internalApi := huma.NewGroup(api, "/internal")
	sessionApi := huma.NewGroup(internalApi, "/session")
	moderatorApi := huma.NewGroup(internalApi, "/moderator")
	adminApi := huma.NewGroup(internalApi, "/admin")

	requireUserSessionMiddlewares = huma.Middlewares{UserAuthHandler, CreateRequireUserAuthHandler(internalApi)}
	requireUserModeratorMiddlewares = huma.Middlewares{UserAuthHandler, CreateRequireUserModeratorHandler(moderatorApi)}
	requireUserAdminMiddlewares = huma.Middlewares{UserAuthHandler, CreateRequireUserAdminHandler(adminApi)}

	moderatorApi.UseMiddleware(requireUserModeratorMiddlewares...)
	adminApi.UseMiddleware(requireUserAdminMiddlewares...)

	registerHealthCheck(internalApi)
	registerAuth(sessionApi)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players",
		OperationID: "get-player",
		Summary:     "Get the current session's Player",
		Description: "get the current session's player",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandleGetSelfPlayer)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/{id}",
		OperationID: "get-player-by-id",
		Summary:     "Get a Player",
		Description: "get a player by ID",
		Tags:        []string{"Player"},
	}, HandleGetPlayer)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/all",
		OperationID: "get-all-players",
		Summary:     "Get all Players",
		Description: "get all players",
		Tags:        []string{"Player"},
	}, HandleGetAllPlayers)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/profile/{id}",
		OperationID: "get-player-profile-by-id",
		Summary:     "Get a Player Profile",
		Description: "get info for a player's profile by ID",
		Tags:        []string{"Player"},
	}, HandleGetPlayerProfile)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/preferredclass/{class}",
		OperationID: "set-player-preferredclass",
		Summary:     "Set Player's preferred class",
		Description: "set the current session player's preferred class by class name",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfPreferredClass)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/preferredlauncher/{launcher}",
		OperationID: "set-player-preferredlauncher",
		Summary:     "Set Player's preferred rocket launcher",
		Description: "set the current session player's preferred rocket launcher by launcher name",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfPreferredLauncher)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/tempusid/{tempus_id}",
		OperationID: "set-player-tempusid",
		Summary:     "Set a Player's own Tempus ID",
		Description: "set a player's own Tempus ID, found at https://tempus2.xyz",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfTempusID)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/steamtradetoken/{url}",
		OperationID: "set-player-steam-trade-token",
		Summary:     "Set a Player's own Steam trade token",
		Description: "set a player's own Steam trade token from their Steam Trade URL, found at https://steamcommunity.com/id/{steamid}/tradeoffers/privacy",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfSteamTradeToken)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/steamavatarurl",
		OperationID: "update-player-steam-avatar-url",
		Summary:     "Update a Player's own Steam avatar url",
		Description: "update a player's own Steam avatar from their Steam Profile",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfSteamAvatarUrl)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/all",
		OperationID: "get-all-full-players",
		Summary:     "Get full info on all Players",
		Description: "get full info on all players",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandleGetAllFullPlayers)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/displayname/{id}/{name}",
		OperationID: "set-player-displayname",
		Summary:     "Set a Player's display name",
		Description: "set a player's display name",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePutSelfDisplayName)
}

// A readiness endpoint is important - it can be used to inform your infrastructure
// (e.g. fly.io) that the API is available. Readiness checks can help keep your API
// alive, by informing fly on when it should try restarting a machine in case of a
// crash.
func registerHealthCheck(internalApi *huma.Group) {
	type ReadyResponse struct{ OK bool }

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/readyz",
		OperationID: "readyz",
		Summary:     "Get Readiness",
		Description: "Get whether or not the API is ready to process requests",
		Tags:        []string{"Health Check"},
	}, func(ctx context.Context, _ *struct{}) (*ReadyResponse, error) {
		return &ReadyResponse{OK: true}, nil
	})
}
