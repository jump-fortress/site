package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	database "github.com/spiritov/jump/api/db/queries"
)

var (
	queries *database.Queries
)

type PlayerIDInput struct {
	ID int64 `path:"id" minimum:"1" doc:"player ID"`
}

type PlayerOutput struct {
	Body database.Player
}

func GetPlayer(ctx context.Context, input *PlayerIDInput) (*PlayerOutput, error) {
	player, err := queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	resp := &PlayerOutput{Body: player}
	return resp, nil
}

func main() {

	db, err := sql.Open("sqlite3", "db/jump.db")
	if err != nil {
		log.Fatalf("[fatal] failed to open db: %v", err)
	}
	defer db.Close()

	queries = database.New(db)

	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("api", "1.0.0"))

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/{id}",
		Summary:     "Get Player",
		Description: "Get a Player by ID",
		Tags:        []string{"Player"},
	}, GetPlayer)

	http.ListenAndServe("localhost:8000", router)
}
