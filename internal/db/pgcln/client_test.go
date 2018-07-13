package pgcln

import (
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// Test_SelectBillsByTime checks if SelectBillsByTime() returns bills from non-empty db
func Test_SelectBillsByTime(t *testing.T) {
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
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgcln.Close()

	start, err := time.Parse(time.RFC3339, "2000-01-01T00:00:00-07:00")
	if err != nil {
		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
	}

	end := time.Now()

	bills, err := pgcln.SelectBillsByTime(start, end)
	if err != nil {
		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
	}

	if len(bills) == 0 {
		t.Fatalf("%v: no bills selected", pgcLogPref)
	}
}

// Test_InsertBill_deleteLastBill tests inserting bill into db, selecting it by service and deleting it
func Test_InsertBill_deleteLastBill(t *testing.T) {
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
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgcln.Close()

	testBill := ServiceBill{
		ID:           1,
		LineItem:     "testItem",
		StartTime:    time.Now(),
		EndTime:      time.Now(),
		Cost:         123.456,
		Currency:     "testCurrency",
		ProjectID:    "testProject",
		Description:  "testDescription",
		GcpCsvFileID: 1,
	}

	if err := pgcln.InsertBill(testBill); err != nil {
		t.Fatalf("%v: insert bill err: %v", pgcLogPref, err)
	}

	bills, err := pgcln.SelectBillsByService("testItem")
	if err != nil {
		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
	}

	if len(bills) != 1 {
		t.Fatalf("%v: incorrect selection, expected 1 bill", pgcLogPref)
	}

	if strings.Compare(bills[0].LineItem, "testItem") != 0 && strings.Compare(bills[0].ProjectID, "testProject") != 0 {
		t.Fatalf("%v: selected bill doesn't match the test bill", pgcLogPref)
	}

	if err := pgcln.deleteLastBill(); err != nil {
		t.Fatalf("%v: delete last bill err: %v", pgcLogPref, err)
	}

	bills, err = pgcln.SelectBillsByService("testItem")
	if err != nil {
		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
	}

	if len(bills) != 0 {
		t.Fatalf("%v: incorrect selection, expected 0 bills", pgcLogPref)
	}
}

// Test_SelectLastBill checks if SelectLastBill() returns the last bill based on test data from db
func Test_SelectLastBill(t *testing.T) {
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
		t.Fatalf("%v: new client err, %v", pgcLogPref, err)
	}
	defer pgcln.Close()

	bill, err := pgcln.SelectLastBill()
	if err != nil {
		t.Fatalf("%v: select bills err: %v", pgcLogPref, err)
	}

	if strings.Compare(bill.ProjectID, "test_project") != 0 {
		t.Fatalf("%v: incorrect selection", pgcLogPref)
	}
}
