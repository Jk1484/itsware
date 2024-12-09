package db

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Pool *pgxpool.Pool

func InitDB() {
	var err error
	connectionString := "host=database user=postgres password=q123 dbname=itsware sslmode=disable"
	Pool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatalln("b5", err)
	}

	err = runMigrations(Pool)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v\n", err)
	}
}

func runMigrations(pool *pgxpool.Pool) error {
	db := stdlib.OpenDBFromPool(pool)

	err := db.Ping()
	if err != nil {
		panic(fmt.Errorf("b:%v", err))
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("1: %v\n", err)
	}

	migrationsPath := fmt.Sprintf("file:%s", "./migrations") //export MIGRATIONS_DIR="E:/path/to/migrationDir"
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"itsware",
		driver,
	)
	if err != nil {
		return fmt.Errorf("2: %v\n", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("3: %v\n", err)
	}

	log.Println("Migrations applied successfully!")
	return nil
}
