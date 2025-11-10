package internal

import (
	"context"
	"log"
	"net/http"

	huma "github.com/danielgtaylor/huma/v2"
	humachi "github.com/danielgtaylor/huma/v2/adapters/humachi"
	chi "github.com/go-chi/chi/v5"
	cors "github.com/rs/cors"
	responses "github.com/spiritov/jump/api/db/responses"
)

var (
	api huma.API
)

func setupRouter() *chi.Mux {
	router := chi.NewMux()

	// spiritov - todo: use strict `AllowedOrigins`
	// spiritov - todo: use CSRF middleware (aa)
	// spiritov - todo: rate limit
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
	registerHealthCheck()

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/{id}",
		Summary:     "Get Player",
		Description: "Get a Player by ID",
		Tags:        []string{"Player"},
	}, responses.GetPlayer)
}

// A readiness endpoint is important - it can be used to inform your infrastructure
// (e.g. fly.io) that the API is available. Readiness checks can help keep your API
// alive, by informing fly on when it should try restarting a machine in case of a
// crash.
func registerHealthCheck() {
	type ReadyResponse struct{ OK bool }

	huma.Register(api, huma.Operation{
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
