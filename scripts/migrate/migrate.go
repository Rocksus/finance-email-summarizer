package main

import (
	"github.com/Rocksus/fundtract/internal/common/config"
	"github.com/Rocksus/fundtract/internal/platform/storage/postgres"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	env := config.GetEnv()
	cfg, err := config.Init("files/etc/fundtract/", env)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed initializing config")
	}

	log.Info().Msg("Connecting to DB...")

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Err(err).Msg("Failed initializing DB driver")
	}

	db := postgres.New(cfg.Database)
	dbGoose := stdlib.OpenDBFromPool(db.FundtractDB.(*pgxpool.Pool))

	if err := goose.Up(dbGoose, "db/migrations/"); err != nil {
		log.Fatal().Err(err).Msg("Failed running DB migration")
	}
}
