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
	pgcLogPref          = "postgres client"
	serviceBillsColumns = 9
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

// Account contains information about GCP user account
type Account struct {
	ID             int
	GcpAccountInfo string
}

// GcpCsvFile contains information about CSV files with billing reports
type GcpCsvFile struct {
	ID           int
	LineItem     string
	StartTime    time.Time
	EndTime      time.Time
	Cost         float64
	Currency     string
	ProjectID    string
	Description  string
	GcpCsvFileID int
}

// ServiceBill contains relevant information from billing report
type ServiceBill struct {
	ID           int
	LineItem     string
	StartTime    time.Time
	EndTime      time.Time
	Cost         float64
	Currency     string
	ProjectID    string
	Description  string
	GcpCsvFileID int
}

// ServiceBills is a set of ServiceBill
type ServiceBills []*ServiceBill

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
		log.Printf("%v: db open err, %v", pgcLogPref, err)
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

// SelectBillsByTime returns bills from db that belong to specified time period
func (c *Client) SelectBillsByTime(start, end time.Time) (ServiceBills, error) {
	if start.After(end) || end.Before(start) {
		return nil, fmt.Errorf("%v: invalid arguments err", pgcLogPref)
	}

	rows, err := c.idb.Query("SELECT * FROM xproject.service_bills ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Printf("%v: db columns err, %v", pgcLogPref, err)
		return nil, err
	}

	if len(cols) != serviceBillsColumns {
		return nil, fmt.Errorf("%v: db format doesn't match ServiceBill struct", pgcLogPref)
	}

	var table ServiceBills
	var row ServiceBill

	result := make([]string, len(cols))
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		for i, raw := range rawResult {
			if raw != nil {
				result[i] = string(raw)
			} else {
				result[i] = "\\N"
			}
		}

		row.ID, err = strconv.Atoi(result[0])
		if err != nil {
			log.Printf("%v: db parse int err, %v", pgcLogPref, err)
			return nil, err
		}

		row.LineItem = result[1]

		row.StartTime, err = time.Parse(time.RFC3339, result[2])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pgcLogPref, err)
			return nil, err
		}

		row.EndTime, err = time.Parse(time.RFC3339, result[3])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pgcLogPref, err)
			return nil, err
		}

		if row.StartTime.Before(start) || row.EndTime.After(end) {
			continue
		}

		row.Cost, err = strconv.ParseFloat(result[4], 64)
		if err != nil {
			log.Printf("%v: db parse float err, %v", pgcLogPref, err)
			return nil, err
		}

		row.Currency = result[5]
		row.ProjectID = result[6]
		row.Description = result[7]
		row.GcpCsvFileID, err = strconv.Atoi(result[8])
		if err != nil {
			log.Printf("%v: db parse int err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// SelectBillsByService returns bills from db that are related to specified GCP service
// If service is an empty string then all bills will be returned
func (c *Client) SelectBillsByService(service string) (ServiceBills, error) {
	rows, err := c.idb.Query("SELECT * FROM xproject.service_bills WHERE line_item LIKE '%" + service + "%' ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Printf("%v: db columns err, %v", pgcLogPref, err)
		return nil, err
	}

	if len(cols) != serviceBillsColumns {
		return nil, fmt.Errorf("%v: db format doesn't match ServiceBill struct", pgcLogPref)
	}

	var table ServiceBills
	var row ServiceBill

	result := make([]string, len(cols))
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		for i, raw := range rawResult {
			if raw != nil {
				result[i] = string(raw)
			} else {
				result[i] = "\\N"
			}
		}

		row.ID, err = strconv.Atoi(result[0])
		if err != nil {
			log.Printf("%v: db parse int err, %v", pgcLogPref, err)
			return nil, err
		}

		row.LineItem = result[1]

		row.StartTime, err = time.Parse(time.RFC3339, result[2])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pgcLogPref, err)
			return nil, err
		}

		row.EndTime, err = time.Parse(time.RFC3339, result[3])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pgcLogPref, err)
			return nil, err
		}

		row.Cost, err = strconv.ParseFloat(result[4], 64)
		if err != nil {
			log.Printf("%v: db parse float err, %v", pgcLogPref, err)
			return nil, err
		}

		row.Currency = result[5]
		row.ProjectID = result[6]
		row.Description = result[7]
		row.GcpCsvFileID, err = strconv.Atoi(result[8])
		if err != nil {
			log.Printf("%v: db parse int err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// SelectLastBill returns the latest added bill from db
func (c *Client) SelectLastBill() (ServiceBill, error) {
	var row ServiceBill

	rows, err := c.idb.Query("SELECT * FROM xproject.service_bills WHERE id = (SELECT MAX(id) FROM xproject.service_bills)")
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return row, err
	}

	cols, err := rows.Columns()
	if err != nil {
		log.Printf("%v: db columns err, %v", pgcLogPref, err)
		return row, err
	}

	if len(cols) != serviceBillsColumns {
		return row, fmt.Errorf("%v: db format doesn't match ServiceBill struct", pgcLogPref)
	}

	result := make([]string, len(cols))
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		if err := rows.Scan(dest...); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return row, err
		}

		for i, raw := range rawResult {
			if raw != nil {
				result[i] = string(raw)
			} else {
				result[i] = "\\N"
			}
		}

		row.ID, err = strconv.Atoi(result[0])
		if err != nil {
			log.Printf("%v: db parse int err, %v", pgcLogPref, err)
			return row, err
		}

		row.LineItem = result[1]

		row.StartTime, err = time.Parse(time.RFC3339, result[2])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pgcLogPref, err)
			return row, err
		}

		row.EndTime, err = time.Parse(time.RFC3339, result[3])
		if err != nil {
			log.Printf("%v: db time parse err, %v", pgcLogPref, err)
			return row, err
		}

		row.Cost, err = strconv.ParseFloat(result[4], 64)
		if err != nil {
			log.Printf("%v: db parse float err, %v", pgcLogPref, err)
			return row, err
		}

		row.Currency = result[5]
		row.ProjectID = result[6]
		row.Description = result[7]
		row.GcpCsvFileID, err = strconv.Atoi(result[8])
		if err != nil {
			log.Printf("%v: db parse int err, %v", pgcLogPref, err)
			return row, err
		}
	}

	return row, nil
}

// InsertBill inserts bill into db
func (c *Client) InsertBill(bill ServiceBill) error {
	if _, err := c.idb.Query("INSERT INTO xproject.service_bills VALUES(DEFAULT, '" +
		bill.LineItem + "', TIMESTAMP '" +
		bill.StartTime.Format(time.RFC3339) + "', TIMESTAMP '" +
		bill.EndTime.Format(time.RFC3339) + "', " +
		strconv.FormatFloat(bill.Cost, 'f', 6, 64) + ", '" +
		bill.Currency + "', '" +
		bill.ProjectID + "', '" +
		bill.Description + "', " +
		strconv.Itoa(bill.GcpCsvFileID) + ")"); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// deleteLastBill deletes the latest added bill from db
func (c *Client) deleteLastBill() error {
	if _, err := c.idb.Query("DELETE FROM xproject.service_bills WHERE id = (SELECT MAX(id) FROM xproject.service_bills)"); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}
