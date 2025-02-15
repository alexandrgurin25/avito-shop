package database

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:embed migrations/*.sql
var fs embed.FS

type ConnConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (c ConnConfig) ToString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
}

func New(connConfig ConnConfig) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(connConfig.ToString())

	if err != nil {
		panic(fmt.Errorf("could not create ConnConfig: %v", err))
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(fmt.Errorf("unable to connect to the database: %v", err))
	}

	defer log.Printf("Postgres connected successful")

	migrations, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		migrations,
		connConfig.ToString(),
	)
	if err != nil {
		log.Printf("Error -> %v", err)
		return nil, fmt.Errorf("unable to migrate the database: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Printf("Unable to migrate the database: %v\n", err)
	}

	return conn, nil
}

type DataBase interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
}

func WithConn() ConnConfig {
	password := os.Getenv("DB_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("DB_PORT can not convert to int", password)
	}
	if password == "" {
		log.Fatal("DB_PASSWORD is not set")
	}

	return ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: password,
		Database: os.Getenv("DB_NAME"),
	}
}
