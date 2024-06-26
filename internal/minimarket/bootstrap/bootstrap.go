package bootstrap

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/nevano11/minimarket/internal/minimarket/config"
	"github.com/nevano11/minimarket/internal/minimarket/handler"
	"github.com/nevano11/minimarket/internal/minimarket/handler/middleware"
	"github.com/nevano11/minimarket/internal/minimarket/repository/postgres"
	"github.com/nevano11/minimarket/internal/minimarket/repository/postgres/repository"
	"github.com/nevano11/minimarket/internal/minimarket/server"
	"github.com/nevano11/minimarket/internal/minimarket/service"
	"github.com/nevano11/minimarket/internal/minimarket/service/auth"
)

func Run(ctx context.Context, cfg Config) error {
	appConfig, err := loadConfig(cfg.ConfigPath)
	if err != nil {
		return err
	}

	db, err := postgres.NewPostgresDb(appConfig.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	storage := repository.NewRepository(db)

	authService := auth.New(storage)
	svc := service.NewService(storage, authService)

	httpHandler := handler.NewHttpHandler(svc, authService, authService, middleware.New(authService))

	routes := httpHandler.InitRoutes()

	apiserver, err := server.NewHttpServer(
		cfg.LogLevel,
		appConfig.Server.Port,
		routes)
	if err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}

	defer apiserver.Shutdown(ctx)

	return apiserver.Start()
}

func loadConfig(configPath string) (*config.Config, error) {
	appConfig, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(appConfig)
	if err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}

	return appConfig, nil
}
