package main

import (
	_ "embed"
	"log"

	"github.com/spiritov/jump/api/db/responses"
	"github.com/spiritov/jump/api/env"
	"github.com/spiritov/jump/api/internal"
	"github.com/spiritov/jump/api/slog"
)

func main() {
	// todo: purpose of providing key JUMP_ENV?
	// env is a wrapper around the `godotenv` library
	if err := env.Load("JUMP_ENV"); err != nil {
		log.Fatalf("[fatal] error loading .env: %v", err)
	}

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

	if err := slog.Setup(); err != nil {
		log.Fatal(err)
	}

	database := responses.OpenDB("./db/jump.db")
	defer database.Close()

	log.Print("db uppies")

	address := env.GetString("JUMP_HTTP_ADDRESS")
	internal.ServeAPI(address)
}
