package main

import (
	_ "embed"
	"log"

	db "github.com/spiritov/jump/api/db"
	env "github.com/spiritov/jump/api/env"
	internal "github.com/spiritov/jump/api/internal"
)

func main() {
	// spiritov - todo: purpose of providing key JUMP_ENV?
	// env is a wrapper around the `godotenv` library
	if err := env.Load("JUMP_ENV"); err != nil {
		log.Fatalf("[fatal] error loading .env: %v", err)
	}

	// spiritov - todo: what's required?
	// spiritov - todo: move required env strings to separate variable / place?
	env.Require(
		"JUMP_SLOG_LEVEL",
		"JUMP_SLOG_MODE",
		"JUMP_HTTP_ADDRESS",
		"JUMP_HTTPLOG_LEVEL",
		"JUMP_HTTPLOG_MODE",
		"JUMP_HTTPLOG_CONCISE",
		"JUMP_HTTPLOG_REQUEST_HEADERS",
		"JUMP_HTTPLOG_RESPONSE_HEADERS",
		"JUMP_HTTPLOG_REQUEST_BODIES",
		"JUMP_HTTPLOG_RESPONSE_BODIES",
		"JUMP_SESSION_TOKEN_SECRET",
		"JUMP_SESSION_COOKIE_SECURE",
		"JUMP_STEAM_API_KEY",
		"JUMP_OID_REALM",
	)

	database := db.OpenDB("api/db/jump.db")
	defer database.Close()

	log.Print("db uppies")

	address := env.GetString("JUMP_HTTP_ADDRESS")
	internal.ServeAPI(address)
}
