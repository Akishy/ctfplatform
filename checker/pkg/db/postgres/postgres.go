package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"sync"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type PGsql struct {
	DB     *pgxpool.Pool
	Logger *zap.Logger
}

func NewPostgresDB(ctx context.Context, logger *zap.Logger, dbConfig *DBConfig) (*PGsql, error) {
	var pgonce sync.Once
	var database *PGsql
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Database)

	pgonce.Do(func() {
		db, err := pgxpool.New(ctx, connStr)
		if err != nil {
			logger.Fatal("Failed to connect to database", zap.Error(err))
		}
		database.DB = db
	})
	database.Logger = logger

	return database, nil
}

func (db *PGsql) Ping() error {
	err := db.DB.Ping(context.Background())
	if err != nil {
		db.Logger.Error("Failed to ping database", zap.Error(err))
		return err
	}
	return nil
}

func (db *PGsql) Close() {
	db.DB.Close()
}
