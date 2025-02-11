package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"strconv"
)

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
	log.Printf("Postgres connected successful")

	defer conn.Close(context.Background()) // Возможно нужно убрать
	return conn, nil
}

type DataBase interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	//ExecContext(ctx context.Context, query string, args ...any) error
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
