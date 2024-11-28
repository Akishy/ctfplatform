package db

import (
	"database/sql"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/config"
	"log"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.PostgresURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

var Module = fx.Provide(NewDB)
