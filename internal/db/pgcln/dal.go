package pgcln

import (
	"context"
	"database/sql"
	"log"
	"time"
)

// selectFromAccounts selects all rows from accounts table
func (c *Client) selectFromAccounts() (*sql.Rows, error) {
	stmt, err := c.idb.PrepareContext(context.Background(), "SELECT * FROM xproject.accounts ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: select from accounts err, %v", pgcLogPref, err)
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(context.Background())
}

// insertIntoAccounts inserts an account into table
func (c *Client) insertIntoAccounts(account GcpAccount) error {
	_, err := c.idb.ExecContext(context.Background(), "INSERT INTO xproject.accounts VALUES(DEFAULT, $1)", account.GcpAccountInfo)

	return err
}

// deleteFromAccounts deletes account with maximum id (latest added account) from table
func (c *Client) deleteFromAccounts() error {
	stmt, err := c.idb.PrepareContext(context.Background(), "DELETE FROM xproject.accounts WHERE id = (SELECT MAX(id) FROM xproject.accounts)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background())

	return err
}

// selectCsvFiles selects all rows from CSV files table
func (c *Client) selectFromCsvFiles() (*sql.Rows, error) {
	stmt, err := c.idb.PrepareContext(context.Background(), "SELECT * FROM xproject.gcp_csv_files ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: select from csv files err, %v", pgcLogPref, err)
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(context.Background())
}

// insertIntoCsvFiles inserts a CSV file into table
func (c *Client) insertIntoCsvFiles(file GcpCsvFile) error {
	_, err := c.idb.ExecContext(context.Background(), "INSERT INTO xproject.gcp_csv_files VALUES(DEFAULT, $1, $2, $3, $4)", file.Name, file.Bucket, file.TimeCreated, file.AccountID)

	return err
}

// deleteFromCsvFiles deletes CSV file with maximum id (latest added file) from table
func (c *Client) deleteFromCsvFiles() error {
	stmt, err := c.idb.PrepareContext(context.Background(), "DELETE FROM xproject.gcp_csv_files WHERE id = (SELECT MAX(id) FROM xproject.gcp_csv_files)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background())

	return err
}

// selectFromBills selects all rows from service bills table
func (c *Client) selectFromBills() (*sql.Rows, error) {
	stmt, err := c.idb.PrepareContext(context.Background(), "SELECT * FROM xproject.service_bills ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(context.Background())
}

// selectBillsByTime selects bills that are within the specified time period
func (c *Client) selectBillsByTime(start, end time.Time) (*sql.Rows, error) {
	return c.idb.QueryContext(context.Background(), "SELECT * FROM xproject.service_bills WHERE start_time >= $1 AND end_time <= $2 ORDER BY id ASC", start, end)
}

// selectBillsByService selects bills that match the specified service
func (c *Client) selectBillsByService(service string) (*sql.Rows, error) {
	service = "%" + service + "%"

	return c.idb.QueryContext(context.Background(), "SELECT * FROM xproject.service_bills WHERE line_item LIKE $1 ORDER BY id ASC", service)
}

// selectBillsByProject selects bills that match the specified project
func (c *Client) selectBillsByProject(project string) (*sql.Rows, error) {
	project = "%" + project + "%"

	return c.idb.QueryContext(context.Background(), "SELECT * FROM xproject.service_bills WHERE project_id LIKE $1 ORDER BY id ASC", project)
}

// selectLastBill selects bill by largest value in end_time, start_time and then id
func (c *Client) selectLastBill() (*sql.Rows, error) {
	stmt, err := c.idb.PrepareContext(context.Background(), "SELECT * FROM xproject.service_bills ORDER BY end_time DESC, start_time DESC, id DESC LIMIT 1")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(context.Background())
}

// insertIntoBills inserts a bill into table
func (c *Client) insertIntoBills(bill ServiceBill) error {
	_, err := c.idb.ExecContext(context.Background(), "INSERT INTO xproject.service_bills VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)", bill.LineItem, bill.StartTime, bill.EndTime, bill.Cost, bill.Currency, bill.ProjectID, bill.Description, bill.GcpCsvFileID)

	return err
}

// deleteFromBills deletes bill with maximum id (latest added bill) from table
func (c *Client) deleteFromBills() error {
	stmt, err := c.idb.PrepareContext(context.Background(), "DELETE FROM xproject.service_bills WHERE id = (SELECT MAX(id) FROM xproject.service_bills)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background())

	return err
}
