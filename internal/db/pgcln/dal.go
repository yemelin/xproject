package pgcln

import (
	"context"
	"database/sql"
	"log"
)

func (c *Client) prepareQueries() error {
	c.queries = make(map[string]*sql.Stmt)
	var err error

	c.queries["selectFromAccounts"], err = c.idb.PrepareContext(context.Background(),
		"SELECT * FROM xproject.accounts ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["insertIntoAccounts"], err = c.idb.PrepareContext(context.Background(),
		"INSERT INTO xproject.accounts VALUES(DEFAULT, $1)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["deleteFromAccounts"], err = c.idb.PrepareContext(context.Background(),
		"DELETE FROM xproject.accounts WHERE id = (SELECT MAX(id) FROM xproject.accounts)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["selectFromCsvFiles"], err = c.idb.PrepareContext(context.Background(),
		"SELECT * FROM xproject.gcp_csv_files ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["insertIntoCsvFiles"], err = c.idb.PrepareContext(context.Background(),
		"INSERT INTO xproject.gcp_csv_files VALUES(DEFAULT, $1, $2, $3, $4)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["deleteFromCsvFiles"], err = c.idb.PrepareContext(context.Background(),
		"DELETE FROM xproject.gcp_csv_files WHERE id = (SELECT MAX(id) FROM xproject.gcp_csv_files)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["selectFromBills"], err = c.idb.PrepareContext(context.Background(),
		"SELECT * FROM xproject.service_bills ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["selectBillsByTime"], err = c.idb.PrepareContext(context.Background(),
		"SELECT * FROM xproject.service_bills WHERE start_time >= $1 AND end_time <= $2 ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["selectBillsByService"], err = c.idb.PrepareContext(context.Background(),
		"SELECT * FROM xproject.service_bills WHERE line_item LIKE '%' || $1 || '%' ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["selectBillsByProject"], err = c.idb.PrepareContext(context.Background(),
		"SELECT * FROM xproject.service_bills WHERE project_id LIKE '%' || $1 || '%' ORDER BY id ASC")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["selectLastBill"], err = c.idb.PrepareContext(context.Background(),
		"SELECT * FROM xproject.service_bills ORDER BY end_time DESC, start_time DESC, id DESC LIMIT 1")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["insertIntoBills"], err = c.idb.PrepareContext(context.Background(),
		"INSERT INTO xproject.service_bills VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	c.queries["deleteFromBills"], err = c.idb.PrepareContext(context.Background(),
		"DELETE FROM xproject.service_bills WHERE id = (SELECT MAX(id) FROM xproject.service_bills)")
	if err != nil {
		log.Printf("%v: prepare err, %v", pgcLogPref, err)
		return err
	}

	return err
}
