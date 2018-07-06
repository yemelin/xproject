package pgcln

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
	// Don't forget add driver importing to main
	// _ "github.com/lib/pq"
)

const (
	pkgLogPref = "postgres client"
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

// Report represents the structure of CSV file
// CSV file used as example: https://storage.googleapis.com/churomann-bucket/test-2018-05-23.csv
type Report struct {
	AccountID   string    `csv:"Account ID"`
	LineItem    string    `csv:"Line Item"`
	StartTime   time.Time `csv:"Start Time"`
	EndTime     time.Time `csv:"End Time"`
	Cost        float64   `csv:"Cost"`
	Currency    string    `csv:"Currency"`
	ProjectID   string    `csv:"Project ID"`
	Description string    `csv:"Description"`
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
		log.Printf("%v: db open err, %v", pkgLogPref, err)
		return nil, err
	}

	// set db interface
	c.idb = db

	return c, nil
}

// Close releases db resources
func (c *Client) Close() error {
	return c.idb.Close()
}

// Ping tests ping to db
func (c *Client) Ping() error {
	return c.idb.Ping()
}

// SelectReports returns reports from db that belong to specified time period
func (c *Client) SelectReports(start, end time.Time) ([]Report, error) {
	if start.After(end) || end.Before(start) {
		return nil, fmt.Errorf("%v: invalid arguments err", pkgLogPref)
	}

	rows, err := c.idb.Query("SELECT * FROM xproject.reports")
	if err != nil {
		log.Printf("%v: db query err, %v", pkgLogPref, err)
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Printf("%v: db columns err, %v", pkgLogPref, err)
		return nil, err
	}

	var table []Report
	var row Report

	result := make([]string, len(cols))
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			log.Printf("%v: db scan err, %v", pkgLogPref, err)
			return nil, err
		}

		for i, raw := range rawResult {
			if raw != nil {
				result[i] = string(raw)
			} else {
				result[i] = "\\N"
			}
		}

		// result[0] is unique id which is not a part of the report
		row.AccountID = result[1]
		row.LineItem = result[2]

		row.StartTime, err = time.Parse(time.RFC3339, result[3])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pkgLogPref, err)
			return nil, err
		}

		row.EndTime, err = time.Parse(time.RFC3339, result[4])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pkgLogPref, err)
			return nil, err
		}

		if row.StartTime.After(end) || row.EndTime.Before(start) {
			continue
		}

		row.Cost, err = strconv.ParseFloat(result[5], 64)
		if err != nil {
			log.Printf("%v: db parse float err, %v", pkgLogPref, err)
			return nil, err
		}

		row.Currency = result[6]
		row.ProjectID = result[7]
		row.Description = result[8]

		table = append(table, row)
	}

	return table, nil
}

// InsertReport inserts a report into db
func (c *Client) InsertReport(report Report) error {
	_, err := c.idb.Query("INSERT INTO xproject.reports VALUES(DEFAULT, '" +
		report.AccountID + "', '" +
		report.LineItem + "', '" +
		report.StartTime.Format(time.RFC3339) + "', '" +
		report.EndTime.Format(time.RFC3339) + "', " +
		strconv.FormatFloat(report.Cost, 'f', 6, 64) + ", '" +
		report.Currency + "', '" +
		report.ProjectID + "', '" +
		report.Description + "')")
	if err != nil {
		log.Printf("%v: db query err, %v", pkgLogPref, err)
		return err
	}

	return nil
}

// DeleteLastReport deletes the last report from db
func (c *Client) DeleteLastReport() error {
	_, err := c.idb.Query("DELETE FROM xproject.reports WHERE id = (SELECT MAX(id) FROM xproject.reports)")
	if err != nil {
		log.Printf("%v: db query err, %v", pkgLogPref, err)
		return err
	}

	return nil
}
