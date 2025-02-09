package sqlite

import (
	"database/sql"

	"github.com/Rocksus/fundtract/internal/platform/log"
	"github.com/pressly/goose"
	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	dsnURI := "fundtract.db"

	// Open the database.
	db, err := sql.Open("sqlite", dsnURI)
	if err != nil {
		log.Fatal("error opening DB: ", err)
	}

	// Optionally, ping the DB to verify the connection.
	if err := db.Ping(); err != nil {
		log.Fatal("error connecting to DB: ", err)
	}

	log.Info("executing migrations...")
	goose.SetDialect("sqlite3")

	if err := goose.Up(db, "./db/migrations/"); err != nil {
		log.Fatal("error running migrations: ", err)
	}

	return db
}
