package main

import (
	"context"

	"github.com/jur4ikoff/netviewer/cmd"
	logger "github.com/jur4ikoff/netviewer/internal"
)

func main() {
	logger := logger.NewLogger()
	ctx := logger.WithContext(context.Background())

	err := cmd.Execute(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to execute command")
	}
}
