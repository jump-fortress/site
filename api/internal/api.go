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
	api huma.API
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
	registerHealthCheck(internalApi)
	registerAuth(sessionApi, internalApi)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/{id}",
		Summary:     "Get Player",
		Description: "get a player by ID",
		Tags:        []string{"Player"},
	}, HandleGetPlayer)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/profile/{id}",
		Summary:     "Get Player Profile",
		Description: "get info for a player's profile by ID",
		Tags:        []string{"Player"},
	}, HandleGetPlayerProfile)
}

// A readiness endpoint is important - it can be used to inform your infrastructure
// (e.g. fly.io) that the API is available. Readiness checks can help keep your API
// alive, by informing fly on when it should try restarting a machine in case of a
// crash.
func registerHealthCheck(internalApi *huma.Group) {
	type ReadyResponse struct{ OK bool }

	huma.Register(internalApi, huma.Operation{
		OperationID: "readyz",
		Method:      http.MethodGet,
		Path:        "/readyz",
		Summary:     "Get Readiness",
		Description: "Get whether or not the API is ready to process requests",
		Tags:        []string{"Health Check"},
	}, func(ctx context.Context, _ *struct{}) (*ReadyResponse, error) {
		return &ReadyResponse{OK: true}, nil
	})
}
