package cmd

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/manager"
	"github.com/doxanocap/hitba-service-api/pkg/httpServer"
	"github.com/doxanocap/hitba-service-api/pkg/redis"
	"github.com/doxanocap/pkg/errs"
	"github.com/doxanocap/pkg/lg"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func SetupManager(
	lc fx.Lifecycle,
	db *gorm.DB,
	redisConn *redis.Conn,
	manager *manager.Manager,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			manager.SetCoreDB(db)
			manager.SetCacheConnection(redisConn)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			connection, err := db.DB()
			if err != nil {
				return errs.Wrap("shutting down: gorm: ", err)
			}
			err = connection.Close()
			if err != nil {
				return errs.Wrap("shutting down: db-conn close: ", err)
			}
			if err := redisConn.Close(); err != nil {
				return errs.Wrap("shutting downs: redis: ", err)
			}
			return nil
		},
	})
}

func RunServer(lc fx.Lifecycle, server *httpServer.Server, manager *manager.Manager) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.Run(manager.Processor().REST().Handler().Engine()); err != nil {
					lg.Fatalf("failed to run REST: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			lg.Info("Stopping server...")
			return server.Shutdown(ctx)
		},
	})
}
