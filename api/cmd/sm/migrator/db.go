package migrator

import (
	"database/sql"
	"errors"
	. "github.com/bdarge/api/cmd/sm/config"
	"github.com/bdarge/api/cmd/sm/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	msq "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Migrate()  (*sql.DB, error) {
	db, _ := sql.Open("mysql", Config.DSN+"&multiStatements=true")

	Config.DB, Config.DBErr = gorm.Open(msq.New(msq.Config{
		Conn: db,
	}), &gorm.Config{})

	err := Config.DB.AutoMigrate(&models.User{}, &models.Account{},
		&models.Customer{}, &models.Order{}, &models.Quote{}, &models.OrderItem{}, &models.QuoteItem{})
	if err != nil {
		return nil, err
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://" + Config.MigrationDir,
		Config.Database,
		driver,
	)

	if err != nil {
		log.Println("ERROR => ", err)
		return nil, err
	}

	err = m.Steps(1)

	if errors.Is(err, os.ErrNotExist) || errors.Is(err, migrate.ErrShortLimit{}) || errors.Is(err, migrate.ErrNoChange) {
		return db, nil
	}

	if err != nil {
		log.Println("ERROR => ", err)
		return nil, err
	}
	log.Printf("Applied migrations")

	return db, nil
}

