package connection

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	dsn string
}

func New(dsn string) *Config {
	return &Config{
		dsn: dsn,
	}
}

func (c *Config) Open() *sql.DB {
	conn, err := sql.Open("postgres", c.dsn)
	if err != nil {
		panic(err.Error())
	}

	return conn
}
