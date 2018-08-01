package pgcln

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	// Don't forget add driver importing to main
	// _ "github.com/lib/pq"
)

const (
	pgcLogPref = "postgres client"
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
		log.Printf("%v: db open err, %v", pgcLogPref, err)
		return nil, err
	}

	// set db interface
	c.idb = db

	return c, nil
}

// combineAccounts combines rows of query result into GcpAccounts
func combineAccounts(rows *sql.Rows) (GcpAccounts, error) {
	var table GcpAccounts
	var row GcpAccount

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.GcpAccountInfo); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// combineCsvFiles combines rows of query result into GcpCsvFiles
func combineCsvFiles(rows *sql.Rows) (GcpCsvFiles, error) {
	var table GcpCsvFiles
	var row GcpCsvFile

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.Name, &row.Bucket, &row.TimeCreated, &row.AccountID); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// combineBills combines rows of query result into ServiceBills
func combineBills(rows *sql.Rows) (ServiceBills, error) {
	var table ServiceBills
	var row ServiceBill

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.LineItem, &row.StartTime, &row.EndTime, &row.Cost, &row.Currency, &row.ProjectID, &row.Description, &row.GcpCsvFileID); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return nil, err
		}

		table = append(table, &row)
	}

	return table, nil
}

// Close releases db resources
func (c *Client) Close() error {
	return c.idb.Close()
}

// Ping tests ping to db
func (c *Client) Ping() error {
	return c.idb.Ping()
}

// ListAccounts returns all accounts from db
func (c *Client) ListAccounts() (GcpAccounts, error) {
	rows, err := c.idb.Query(selectFrom("xproject.accounts"))
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineAccounts(rows)
}

// AddAccount adds account into db
func (c *Client) AddAccount(account GcpAccount) error {
	if _, err := c.idb.Query(insertInto("xproject.accounts", 1), account.GcpAccountInfo); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// removeLastAccount removes the latest added account from db
func (c *Client) removeLastAccount() error {
	if _, err := c.idb.Query(deleteFrom("xproject.accounts")); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// ListCsvFiles returns all CSV files from db
func (c *Client) ListCsvFiles() (GcpCsvFiles, error) {
	rows, err := c.idb.Query(selectFrom("xproject.gcp_csv_files"))
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineCsvFiles(rows)
}

// AddCsvFile adds CSV file into db
func (c *Client) AddCsvFile(file GcpCsvFile) error {
	if _, err := c.idb.Query(insertInto("xproject.gcp_csv_files", 4), file.Name, file.Bucket, file.TimeCreated, file.AccountID); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// removeLastCsvFile removes the latest added CSV file from db
func (c *Client) removeLastCsvFile() error {
	if _, err := c.idb.Query(deleteFrom("xproject.gcp_csv_files")); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// ListAllBills returns all bills from db
func (c *Client) ListAllBills() (ServiceBills, error) {
	rows, err := c.idb.Query(selectFrom("xproject.service_bills"))
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// ListBillsByTime returns bills from db that are within the specified time period
func (c *Client) ListBillsByTime(start, end time.Time) (ServiceBills, error) {
	if start.After(end) || end.Before(start) {
		return nil, fmt.Errorf("%v: invalid arguments err", pgcLogPref)
	}

	rows, err := c.idb.Query(selectByRange("xproject.service_bills", "start_time", "end_time"), start, end)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// ListBillsByService returns bills from db that are related to specified GCP service
// If service is an empty string then all bills will be returned
func (c *Client) ListBillsByService(service string) (ServiceBills, error) {
	service = "%" + service + "%"

	rows, err := c.idb.Query(selectByMatch("xproject.service_bills", "line_item"), service)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// ListBillsByProject returns bills from db that are related to specified GCP project
// If project is an empty string then all bills will be returned
func (c *Client) ListBillsByProject(project string) (ServiceBills, error) {
	project = "%" + project + "%"

	rows, err := c.idb.Query(selectByMatch("xproject.service_bills", "project_id"), project)
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return nil, err
	}
	defer rows.Close()

	return combineBills(rows)
}

// GetLastBill returns the latest added bill from db by time
func (c *Client) GetLastBill() (ServiceBill, error) {
	var row ServiceBill

	rows, err := c.idb.Query(selectLast("xproject.service_bills", "end_time", "start_time", "id"))
	if err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return row, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&row.ID, &row.LineItem, &row.StartTime, &row.EndTime, &row.Cost, &row.Currency, &row.ProjectID, &row.Description, &row.GcpCsvFileID); err != nil {
			log.Printf("%v: db scan err, %v", pgcLogPref, err)
			return row, err
		}
	}

	return row, nil
}

// AddBill adds bill into db
func (c *Client) AddBill(bill ServiceBill) error {
	if _, err := c.idb.Query(insertInto("xproject.service_bills", 8), bill.LineItem, bill.StartTime, bill.EndTime, bill.Cost, bill.Currency, bill.ProjectID, bill.Description, bill.GcpCsvFileID); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}

// removeLastBill removes the latest added bill from db
func (c *Client) removeLastBill() error {
	if _, err := c.idb.Query(deleteFrom("xproject.service_bills")); err != nil {
		log.Printf("%v: db query err, %v", pgcLogPref, err)
		return err
	}

	return nil
}
