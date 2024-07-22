package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

var (
	DbConnection *sql.DB
)

func Initialize(dbParam *sql.DB) {
	DbConnection = dbParam
}

// func DbMigrate(dbParam *sql.DB) {
// 	migrations := &migrate.PackrMigrationSource{
// 		Box: packr.New("migrations", "./sql_migrations"),
// 	}

// 	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
// 	if errs != nil {
// 		panic(errs)
// 	}

// 	DbConnection = dbParam

// 	fmt.Println("Applied: ", n, "migrations!")
// }

//go:embed sql_migrations/*.sql
var migrationsFS embed.FS

func DbMigrate(dbParam *sql.DB) {
	// Create a new instance of the Postgres driver
	driver, err := postgres.WithInstance(dbParam, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	// Create a new source driver for the embedded files
	source, err := iofs.New(migrationsFS, "sql_migrations")
	if err != nil {
		panic(err)
	}

	// Create a new instance of migrate with the Postgres driver and the embedded source
	m, err := migrate.NewWithInstance("iofs", source, "postgres", driver)
	if err != nil {
		panic(err)
	}

	// Apply the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	// Set the global database connection
	DbConnection = dbParam

	fmt.Println("Migrations applied successfully!")
}
