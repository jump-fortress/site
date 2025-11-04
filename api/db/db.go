package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	database "github.com/spiritov/jump/api/db/queries"
)

var (
	queries *database.Queries
)

func OpenDB(path string) *sql.DB {
	// os.Remove("./db/jump.db")
	db, err := sql.Open("sqlite3", "./db/jump.db")
	if err != nil {
		log.Fatalf("[fatal] failed to open db: %v", err)
	}

	queries = database.New(db)
	return db
}

func GetPlayer(ctx context.Context, input *PlayerIDInput) (*PlayerResponse, error) {
	player, err := queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	resp := &PlayerResponse{Body: player}
	return resp, nil
}
