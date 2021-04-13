package httputil

import (
	"database/sql"
	"errors"
	. "github.com/bdarge/sm-api/cmd/invoice/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func Migrate()  (*sql.DB, error) {
	db, _ := sql.Open("mysql", Config.DSN+"&multiStatements=true")

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"profile",
		driver,
	)

	err := m.Steps(1)

	if errors.Is(err, os.ErrNotExist) || errors.Is(err, migrate.ErrShortLimit{})  || errors.Is(err, migrate.ErrNoChange) {
		return db, nil
	}

	if err != nil {
		log.Println("ERROR => ", err)
		return nil, err
	}
	log.Printf("Applied migrations")

	return db, nil
}
