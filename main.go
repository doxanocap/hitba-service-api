package main

import (
	"github.com/doxanocap/hitba-service-api/cmd"
	"github.com/doxanocap/hitba-service-api/internal/config"
	"github.com/doxanocap/hitba-service-api/internal/manager"
	"github.com/doxanocap/hitba-service-api/pkg/banner"
	"github.com/doxanocap/hitba-service-api/pkg/httpServer"
	"github.com/doxanocap/hitba-service-api/pkg/logger"
	"github.com/doxanocap/hitba-service-api/pkg/postgres"
	"github.com/doxanocap/hitba-service-api/pkg/redis"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			config.InitConfig,
			logger.InitLogger,
			postgres.InitConnection,
			redis.InitConnection,
			manager.InitManager,
			httpServer.InitServer,
		),
		fx.Invoke(
			cmd.SetupManager,
			cmd.RunServer,
			banner.Default,
		),
	)

	app.Run()
	if err := app.Err(); err != nil {
		logger.Log.Fatal(err.Error())
	}
}
