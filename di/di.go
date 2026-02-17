package di

import (
	"chess-console/configs"
	"chess-console/internal/games"
	"log/slog"
)

type Container struct {
	Config    *configs.Config
	Games     games.Games
	Validator *CustomValidator
	Logger    *slog.Logger
}

func SetUp() *Container {
	var (
		cfg       = configs.LoadDefault()
		validator = NewCustomValidator()
		logger    = newLogger(cfg)
	)

	// Initialize services
	gameService := games.NewGame(logger)

	return &Container{
		Config:    cfg,
		Games:     gameService,
		Validator: validator,
		Logger:    logger,
	}
}
