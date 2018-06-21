// +build integration

package pgcln

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

//
func TestInteg_NewPing(t *testing.T) {

	conf := Config{
		Host:     os.Getenv(EnvDBHost),
		Port:     os.Getenv(EnvDBPort),
		DB:       os.Getenv(EnvDBName),
		User:     os.Getenv(EnvDBUser),
		Password: os.Getenv(EnvDBPwd),
		SSLMode:  "disable",
	}

	pgcln, err := New(conf)
	if err != nil {
		t.Fatalf("%v: new client err, %v", pgkLogPref, err)
	}

	// ping
	err = pgcln.Ping()
	if err != nil {
		t.Fatalf("%v: pg client ping err, %v", pgkLogPref, err)
	}
}
