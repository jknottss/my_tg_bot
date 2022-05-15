package database

import (
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Connection struct {
	Pool *pgx.ConnPool
}

func getConfig() (*pgx.ConnPoolConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	PsqlPort := os.Getenv("PSQL_PORT")
	Port, _ := strconv.ParseUint(PsqlPort, 10, 64)
	conf := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     os.Getenv("PSQL_HOST"),
			Port:     uint16(Port),
			Password: os.Getenv("PSQL_PASSWORD"),
			Database: os.Getenv("PSQL_DB"),
			User:     os.Getenv("PSQL_USER"),
		}}
	return &conf, nil
}

func Connect() (*Connection, error) {
	config, err := getConfig()
	if err != nil {
		return nil, err
	}
	if pool, err := pgx.NewConnPool(*config); err == nil {
		return &Connection{pool}, nil
	} else {
		return nil, err
	}
}
