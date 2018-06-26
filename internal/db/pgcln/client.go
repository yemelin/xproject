package pgcln

import (
	"database/sql"
	"fmt"
	"log"
	// Don't forget add driver importing to main
	// _ "github.com/lib/pq"
)

const (
	pgkLogPref = "postgres client"
)

// Env for testing
const (
	EnvDBHost = "APP_DB_PG_HOST"
	EnvDBPort = "APP_DB_PG_PORT"
	EnvDBName = "APP_DB_PG_NAME"
	EnvDBUser = "APP_DB_PG_USER"
	EnvDBPwd  = "APP_DB_PG_PWD"
)

// Config sets database configs
type Config struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	SSLMode  string
}

// Client implements postgres db client
type Client struct {
	// config
	conf Config

	// db
	idb IDB
}

// New inits client
func New(conf Config) (*Client, error) {
	c := &Client{
		conf: conf,
	}

	// get db
	dbSourceName := fmt.Sprintf(
		"host=%v port=%v dbname=%v user=%v password=%v sslmode=%v",
		conf.Host, conf.Port, conf.DB, conf.User, conf.Password, conf.SSLMode,
	)
	db, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		log.Printf("%v: db open err, %v", pgkLogPref, err)
		return nil, err
	}

	// set db interface
	c.idb = db

	return c, nil
}

// Close releses db resources
func (c *Client) Close() error {
	return c.idb.Close()
}

// Ping test pint to db
func (c *Client) Ping() error {
	return c.idb.Ping()
}
