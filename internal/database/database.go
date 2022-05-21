package database

import (
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const create = `
		CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		user_id TEXT,
		task TEXT,
		priority INT,
		done BOOL
		); `

type Connection struct {
	Pool *pgx.ConnPool
}

func getConfig() (*pgx.ConnPoolConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	PsqlPort := os.Getenv("POSTGRES_PORT")
	Port, _ := strconv.ParseUint(PsqlPort, 10, 64)
	conf := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     uint16(Port),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
			User:     os.Getenv("POSTGRES_USER"),
		}}
	return &conf, nil
}

func Connect() (*Connection, error) {
	config, err := getConfig()
	if err != nil {
		return nil, err
	}
	if pool, err := pgx.NewConnPool(*config); err == nil {
		_, err = pool.Exec(create)
		if err != nil {
			return nil, err
		}
		return &Connection{pool}, nil
	} else {
		return nil, err
	}
}
