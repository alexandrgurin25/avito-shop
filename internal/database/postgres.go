package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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

func (c ConnConfig) New(connConfig ConnConfig) *pgx.Conn {
	config, err := pgx.ParseConfig(c.ToString())

	if err != nil {
		panic(fmt.Errorf("could not create ConnConfig: %v", err))
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(fmt.Errorf("unable to connect to the database: %v", err))
	}
	log.Printf("Postgres connected successful")

	defer conn.Close(context.Background()) // Возможно нужно убрать
	return conn
}

type DataBase interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	ExecContext(ctx context.Context, query string, args ...any) error
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func WithConn() ConnConfig {
	password := os.Getenv("DB_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("DB_PORT can not convert to int")
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
