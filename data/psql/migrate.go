package psql

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/mochahub/coinprice-scraper/config"
	"log"
)

func RunMigrations(
	db *sqlx.DB,
	secrets *config.Secrets,
) (int, error) {
	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{DatabaseName: secrets.DBName})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}
	sourceURL := fmt.Sprintf("file://%s/data/psql/migrations", config.GetProjectRoot())
	m, err := migrate.NewWithDatabaseInstance(
		sourceURL, // file://path/to/directory
		secrets.DatabaseCredentials.DBName, driver)
	if err != nil {
		log.Fatalf("migration failed... %v", err)
		return 0, err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
		return 0, err
	}
	log.Println("Database migrated")
	verion, _, err := m.Version()
	return int(verion), err
}
