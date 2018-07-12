package pgcln

import (
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// Test_SelectReportsByTime checks if SelectReportsByTime() returns reports from non-empty db
func Test_SelectReportsByTime(t *testing.T) {
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

	reports, err := pgcln.SelectReportsByTime(start, end)
	if err != nil {
		t.Fatalf("%v: select reports err: %v", pgcLogPref, err)
	}

	if len(reports) == 0 {
		t.Fatalf("%v: no reports selected", pgcLogPref)
	}
}

// Test_InsertReport_deleteLastReport tests inserting report into db, selecting it by service and deleting it
func Test_InsertReport_deleteLastReport(t *testing.T) {
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

	testReport := Report{
		AccountID:   "testAccount",
		LineItem:    "testItem",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		Cost:        123.456,
		Currency:    "testCurrency",
		ProjectID:   "testProject",
		Description: "testDescription",
	}

	if err := pgcln.InsertReport(testReport); err != nil {
		t.Fatalf("%v: insert report err: %v", pgcLogPref, err)
	}

	reports, err := pgcln.SelectReportsByService("testItem")
	if err != nil {
		t.Fatalf("%v: select reports err: %v", pgcLogPref, err)
	}

	if len(reports) != 1 {
		t.Fatalf("%v: incorrect selection, expected 1 report", pgcLogPref)
	}

	if strings.Compare(reports[0].AccountID, "testAccount") != 0 || strings.Compare(reports[0].ProjectID, "testProject") != 0 {
		t.Fatalf("%v: selected report doesn't match the test report", pgcLogPref)
	}

	if err := pgcln.deleteLastReport(); err != nil {
		t.Fatalf("%v: delete last report err: %v", pgcLogPref, err)
	}

	reports, err = pgcln.SelectReportsByService("testItem")
	if err != nil {
		t.Fatalf("%v: select reports err: %v", pgcLogPref, err)
	}

	if len(reports) != 0 {
		t.Fatalf("%v: incorrect selection, expected 0 reports", pgcLogPref)
	}
}

// Test_SelectLastReport checks if SelectLastReport() returns the last report based on test data from db
func Test_SelectLastReport(t *testing.T) {
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

	report, err := pgcln.SelectLastReport()
	if err != nil {
		t.Fatalf("%v: select reports err: %v", pgcLogPref, err)
	}

	if strings.Compare(report.ProjectID, "proj-2-205013") != 0 {
		t.Fatalf("%v: incorrect selection", pgcLogPref)
	}
}
