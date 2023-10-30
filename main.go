package main

import (
	"github.com/doxanocap/hitba-service-api/cmd"
	_ "github.com/doxanocap/hitba-service-api/docs"
	"github.com/doxanocap/hitba-service-api/internal/config"
	"github.com/doxanocap/hitba-service-api/internal/manager"
	"github.com/doxanocap/hitba-service-api/pkg/banner"
	"github.com/doxanocap/hitba-service-api/pkg/httpServer"
	"github.com/doxanocap/hitba-service-api/pkg/logger"
	"github.com/doxanocap/hitba-service-api/pkg/postgres"
	"github.com/doxanocap/hitba-service-api/pkg/redis"
	"github.com/doxanocap/pkg/lg"
	"go.uber.org/fx"
)

//	@title			service-api
//	@version		1.0
//	@description	API Server for Admin Application

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			config.InitConfig,
			postgres.InitConnection,
			redis.InitConnection,
			manager.InitManager,
			httpServer.InitServer,
		),
		fx.Invoke(
			logger.InitLogger,
			cmd.SetupManager,
			cmd.RunServer,
			banner.Default,
		),
	)

	app.Run()
	if err := app.Err(); err != nil {
		lg.Fatal(err)
	}
}
