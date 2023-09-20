package postgres

import (
	"fmt"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/lg"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	driver        = "postgres"
	ChatListTable = "chat_list"
	ChatMessages  = "chat_messages"
)

func getDSN(cfg model.Psql) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PsqlHost, cfg.PsqlPort, cfg.PsqlUser, cfg.PsqlPassword, cfg.PsqlDatabase, cfg.PsqlSSL)
}

func InitConnection(cfg *model.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(getDSN(cfg.Psql)), &gorm.Config{})
	if err != nil {
		lg.Fatalf("psql: failed to connect -> %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		lg.Fatalf("psql: failed to set connection -> %v", err)
	}

	err = connection.Ping()
	if err != nil {
		lg.Fatalf("psql: failed to ping -> %v", err)
	}

	return db
}
