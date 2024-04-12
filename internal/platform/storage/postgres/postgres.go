package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Rocksus/fundtract/internal/common/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	dbConnectionStringURL = "%s://%s:%s@%s:%d/%s"
)

func New(cfg map[string]*config.DatabaseConfig) *Connections {
	fundtractDB := createConnection(cfg["fundtract"])

	return &Connections{
		FundtractDB: fundtractDB,
	}
}

func createConnection(cfg *config.DatabaseConfig) Pool {
	addr := createConnectionString(cfg)

	dbPoolConfig, err := pgxpool.ParseConfig(addr)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to parse pool config")
	}

	if cfg.MaxOpenConn > 0 {
		dbPoolConfig.MaxConns = int32(cfg.MaxOpenConn)
	}
	if cfg.ConnMaxLifetime > 0 {
		dbPoolConfig.MaxConnLifetime = time.Duration(cfg.ConnMaxLifetime) * time.Second
	} else {
		dbPoolConfig.MaxConnLifetime = 5 * time.Second
	}
	if cfg.ConnMaxIdleTime > 0 {
		dbPoolConfig.MaxConnIdleTime = time.Duration(cfg.ConnMaxIdleTime) * time.Second
	} else {
		dbPoolConfig.MaxConnIdleTime = 30 * time.Minute
	}

	pool, err := ConnectPgxPool(context.Background(), dbPoolConfig)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create pool")
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal().
			Str("db_name", cfg.Name).
			Msg("Can't connect to DB")
	}

	return pool
}

func createConnectionString(cfg *config.DatabaseConfig) string {
	credentialsMaster := strings.Split(cfg.Credentials, ":")

	connectionString := fmt.Sprintf(dbConnectionStringURL,
		cfg.Driver,
		credentialsMaster[0],
		credentialsMaster[1],
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	sslParams := make([]string, 0)

	sslMode := "disable"
	if cfg.SSLRootCert != "" && cfg.SSLCert != "" && cfg.SSLKey != "" {
		sslMode = "verify-full"
	}
	sslParams = append(sslParams, fmt.Sprintf("sslmode=%s", sslMode))

	if cfg.SSLRootCert != "" {
		sslParams = append(sslParams, fmt.Sprintf("sslrootcert=%s", cfg.SSLRootCert))
	}
	if cfg.SSLCert != "" {
		sslParams = append(sslParams, fmt.Sprintf("sslcert=%s", cfg.SSLCert))
	}
	if cfg.SSLKey != "" {
		sslParams = append(sslParams, fmt.Sprintf("sslkey=%s", cfg.SSLKey))
	}

	// Append the SSL parameters to the connection string
	connectionStringWithParams := connectionString + "?" + strings.Join(sslParams, "&")

	return connectionStringWithParams
}
