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

// Report represents the structure of CSV file
// CSV file used as example: https://storage.googleapis.com/churomann-bucket/test-2018-05-23.csv
type Report struct {
	AccountID              string `csv:"Account ID"`
	LineItem               string `csv:"Line Item"`
	StartTime              string `csv:"Start Time"`
	EndTime                string `csv:"End Time"`
	Project                string `csv:"Project"`
	Measurement            string `csv:"Measurement1"`
	MeasurementConsumption string `csv:"Measurement1 Total Consumption"`
	MeasurementUnits       string `csv:"Measurement1 Units"`
	Credit                 string `csv:"Credit1"`
	CreditAmount           string `csv:"Credit1 Amount"`
	CreditCurrency         string `csv:"Credit1 Currency"`
	Cost                   string `csv:"Cost"`
	Currency               string `csv:"Currency"`
	ProjectNumber          string `csv:"Project Number"`
	ProjectID              string `csv:"Project ID"`
	ProjectName            string `csv:"Project Name"`
	ProjectLabels          string `csv:"Project Labels"`
	Description            string `csv:"Description"`
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

// Close releases db resources
func (c *Client) Close() error {
	return c.idb.Close()
}

// Ping tests ping to db
func (c *Client) Ping() error {
	return c.idb.Ping()
}

// SelectReports returns all reports from db
func (c *Client) SelectReports() ([]Report, error) {
	rows, err := c.idb.Query("SELECT * FROM xproject.reports")
	if err != nil {
		log.Printf("%v: db query err, %v", pgkLogPref, err)
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Printf("%v: db columns err, %v", pgkLogPref, err)
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
			log.Printf("%v: db scan err, %v", pgkLogPref, err)
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
		row.StartTime = result[3]
		row.EndTime = result[4]
		row.Project = result[5]
		row.Measurement = result[6]
		row.MeasurementConsumption = result[7]
		row.MeasurementUnits = result[8]
		row.Credit = result[9]
		row.CreditAmount = result[10]
		row.CreditCurrency = result[11]
		row.Cost = result[12]
		row.Currency = result[13]
		row.ProjectNumber = result[14]
		row.ProjectID = result[15]
		row.ProjectName = result[16]
		row.ProjectLabels = result[17]
		row.Description = result[18]

		table = append(table, row)
	}

	return table, nil
}

// InsertReport inserts a report into db
func (c *Client) InsertReport(report Report) error {
	_, err := c.idb.Query("INSERT INTO xproject.reports VALUES(DEFAULT, '" +
		report.AccountID + "', '" +
		report.LineItem + "', '" +
		report.StartTime + "', '" +
		report.EndTime + "', '" +
		report.Project + "', '" +
		report.Measurement + "', '" +
		report.MeasurementConsumption + "', '" +
		report.MeasurementUnits + "', '" +
		report.Credit + "', '" +
		report.CreditAmount + "', '" +
		report.CreditCurrency + "', '" +
		report.Cost + "', '" +
		report.Currency + "', '" +
		report.ProjectNumber + "', '" +
		report.ProjectID + "', '" +
		report.ProjectName + "', '" +
		report.ProjectLabels + "', '" +
		report.Description + "')")
	if err != nil {
		log.Printf("%v: db query err, %v", pgkLogPref, err)
		return err
	}

	return nil
}

// DeleteLastReport deletes the last report from db
func (c *Client) DeleteLastReport() error {
	_, err := c.idb.Query("DELETE FROM xproject.reports WHERE id = (SELECT MAX(id) FROM xproject.reports)")
	if err != nil {
		log.Printf("%v: db query err, %v", pgkLogPref, err)
		return err
	}

	return nil
}
