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
	PsqlHost := os.Getenv("PSQL_HOST")
	PsqlPort := os.Getenv("PSQL_PORT")
	PsqlUser := os.Getenv("PSQL_USER")
	PsqlPassword := os.Getenv("PSQL_PASSWORD")
	PsqlDb := os.Getenv("PSQL_DB")
	Port, _ := strconv.ParseUint(PsqlPort, 10, 64)
	conf := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     PsqlHost,
			Port:     uint16(Port),
			Password: PsqlPassword,
			Database: PsqlDb,
			User:     PsqlUser,
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

func (c *Connection) AddTask(userId int) error {

}

func (c *Connection) DropTask(userId int) error {

}

func (c *Connection) ShowAllTasks(userId int) error {

}
