package slog

import (
	"log/slog"
	"os"

	"github.com/rotisserie/eris"
	"github.com/spiritov/jump/api/env"
)

var Logger *slog.Logger

var SlogLevelMap = map[string]slog.Level{
	"Debug": slog.LevelDebug,
	"Info":  slog.LevelInfo,
	"Warn":  slog.LevelWarn,
	"Error": slog.LevelError,
}

func Setup() error {
	logLevel, matchedErr := env.GetMapped("JUMP_SLOG_LEVEL", SlogLevelMap)
	if matchedErr != nil {
		return matchedErr
	}

	handlerOptions := &slog.HandlerOptions{Level: logLevel}

	slogMode := env.GetString("JUMP_SLOG_MODE")
	switch slogMode {
	case "Text":
		Logger = slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))
	case "JSON":
		Logger = slog.New(slog.NewJSONHandler(os.Stdout, handlerOptions))
	default:
		return eris.Errorf("invalid value for JUMP_SLOG_MODE: %s", slogMode)
	}

	return nil
}
