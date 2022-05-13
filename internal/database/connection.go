package database

import (
	"bot/internal/getconfig"
	"github.com/jackc/pgx"
)

type Connection struct {
	conn *pgx.ConnPool
}

func Connect() (*Connection, error) {
	config, err := getconfig.GetConfig()
	if err != nil {
		return nil, err
	}
	if pool, err := pgx.NewConnPool(*config); err == nil {
		return &Connection{pool}, nil
	} else {
		return nil, err
	}
}
