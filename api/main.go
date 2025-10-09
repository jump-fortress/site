package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
	database "github.com/spiritov/jump/api/db/queries"
)

func main() {
	db, err := sql.Open("sqlite3", "db/jump.db")
	if err != nil {
		log.Fatalf("[fatal] failed to open db: %v", err)
	}
	defer db.Close()

	queries := database.New(db)
	ctx := context.Background()
	player, err := queries.InsertPlayer(ctx, database.InsertPlayerParams{
		SteamID:         "test",
		DisplayName:     "spiritov",
		SoldierDivision: sql.NullString{String: "div", Valid: true},
		DemoDivision:    sql.NullString{},
	})
	if err != nil {
		log.Printf("[error] failed to insert player: %v", err)
	} else {
		log.Printf("got player: %+v", player)
	}

}
