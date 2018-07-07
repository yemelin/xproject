package pgcln

import (
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// Test_SelectReports checks if SelectReports() returns reports from non-empty db
func Test_SelectReports(t *testing.T) {
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

	start, err := time.Parse(time.RFC3339, "2018-05-22T00:00:00-07:00")
	if err != nil {
		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
	}

	end, err := time.Parse(time.RFC3339, "2018-05-23T00:00:00-07:00")
	if err != nil {
		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
	}

	reports, err := pgcln.SelectReports(start, end)
	if err != nil {
		t.Fatalf("%v: select reports err: %v", pgcLogPref, err)
	}

	if len(reports) == 0 {
		t.Fatalf("%v: no reports selected", pgcLogPref)
	}
}

// Test_InsertReport_DeleteLastReport tests inserting report into db and deleting it
func Test_InsertReport_DeleteLastReport(t *testing.T) {
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

	testStart, err := time.Parse(time.RFC3339, "2018-05-26T00:00:00-07:00")
	if err != nil {
		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
	}

	testEnd, err := time.Parse(time.RFC3339, "2018-05-27T00:00:00-07:00")
	if err != nil {
		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
	}

	// Only testReport belongs to this time period
	start, err := time.Parse(time.RFC3339, "2018-05-25T00:00:00-07:00")
	if err != nil {
		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
	}

	end, err := time.Parse(time.RFC3339, "2018-05-28T00:00:00-07:00")
	if err != nil {
		t.Fatalf("%v: time parse err, %v", pgcLogPref, err)
	}

	testReport := Report{
		AccountID:   "testAccount",
		LineItem:    "testItem",
		StartTime:   testStart,
		EndTime:     testEnd,
		Cost:        123.456,
		Currency:    "testCurrency",
		ProjectID:   "testProject",
		Description: "testDescription",
	}

	if err := pgcln.InsertReport(testReport); err != nil {
		t.Fatalf("%v: insert report err: %v", pgcLogPref, err)
	}

	reports, err := pgcln.SelectReports(start, end)
	if err != nil {
		t.Fatalf("%v: select reports err: %v", pgcLogPref, err)
	}

	if len(reports) != 1 {
		t.Fatalf("%v: incorrect selection, expected 1 report", pgcLogPref)
	}

	if strings.Compare(reports[0].AccountID, "testAccount") != 0 || strings.Compare(reports[0].LineItem, "testItem") != 0 {
		t.Fatalf("%v: selected report doesn't match the test report", pgcLogPref)
	}

	if err := pgcln.DeleteLastReport(); err != nil {
		t.Fatalf("%v: delete last report err: %v", pgcLogPref, err)
	}

	reports, err = pgcln.SelectReports(start, end)
	if err != nil {
		t.Fatalf("%v: select reports err: %v", pgcLogPref, err)
	}

	if len(reports) != 0 {
		t.Fatalf("%v: incorrect selection, expected 0 reports", pgcLogPref)
	}
}
