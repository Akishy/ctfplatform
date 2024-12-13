package migrations_module

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	// Init DB drivers.
	_ "github.com/lib/pq"
)

func makeMigrations(connectionString string) {
	db, err := goose.OpenDBWithDriver("pgx", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations_module"); err != nil {
		log.Fatal(err)
	}
}
