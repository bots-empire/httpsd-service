package db

import (
	"context"
	"database/sql"
	"fmt"
	"httpsd-service/db"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

const dbDriver = "postgres"

func InitDataBase(ctx context.Context, cfg *pgxpool.Config) (*pgxpool.Pool, error) {
	dataBase, err := sql.Open(dbDriver, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 4343, "httpsd-user", "root321root", "httpsd-service"))
	if err != nil {
		log.Fatalf("Failed open database: %s\n", err.Error())
	}

	goose.SetBaseFS(db.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(dataBase, "migrations"); err != nil {
		panic(err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create conn pool")
	}

	return pool, nil
}
