package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(c *DBConfig) (*pgx.Conn, error) {
	dbUrl := FormatDBUrl(c)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	defer conn.Close(context.Background())

	return conn, err
}
