package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/spf13/viper"
)

func Runmigration() {
	db, err := sql.Open("postgres", viper.GetString("DB_URL"))
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal("errors to connect", err)
	}
	if err := m.Up(); err != nil {
		//Если ничего не накатилось то не нужно ругаться
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
