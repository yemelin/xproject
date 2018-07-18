package pgcln

import (
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// Test_Account tests adding account into db, listing and removing it
func Test_Account(t *testing.T) {
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

	testAccount := Account{
		ID:             1,
		GcpAccountInfo: "testInfo",
	}

	if err := pgcln.AddAccount(testAccount); err != nil {
		t.Fatalf("%v: add account err: %v", pgcLogPref, err)
	}

	accounts, err := pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	if len(accounts) != 2 {
		t.Fatalf("%v: expected 2 accounts, not %v", pgcLogPref, len(accounts))
	}

	if strings.Compare(accounts[0].GcpAccountInfo, "testInfo") != 0 {
		t.Fatalf("%v: account's info doesn't match the test account", pgcLogPref)
	}

	if err := pgcln.removeLastAccount(); err != nil {
		t.Fatalf("%v: remove last account err: %v", pgcLogPref, err)
	}

	accounts, err = pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	if len(accounts) != 1 {
		t.Fatalf("%v: expected 1 account, not %v", pgcLogPref, len(accounts))
	}
}

// Test_CsvFile tests Adding CSV file into db, Listing and removing it
func Test_CsvFile(t *testing.T) {
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

	testCsvFile := GcpCsvFile{
		ID:          1,
		Name:        "testName",
		Bucket:      "testBucket",
		TimeCreated: time.Now(),
		AccountID:   1,
	}

	if err := pgcln.AddCsvFile(testCsvFile); err != nil {
		t.Fatalf("%v: add csv file err: %v", pgcLogPref, err)
	}

	csvFiles, err := pgcln.ListCsvFiles()
	if err != nil {
		t.Fatalf("%v: list csv file err: %v", pgcLogPref, err)
	}

	if len(csvFiles) != 2 {
		t.Fatalf("%v: expected 2 csv files, not %v", pgcLogPref, len(csvFiles))
	}

	if strings.Compare(csvFiles[0].Name, "testName") != 0 {
		t.Fatalf("%v: csv file's name doesn't match the test csv file", pgcLogPref)
	}

	if err := pgcln.removeLastCsvFile(); err != nil {
		t.Fatalf("%v: remove last csv file err: %v", pgcLogPref, err)
	}

	csvFiles, err = pgcln.ListCsvFiles()
	if err != nil {
		t.Fatalf("%v: list csv files err: %v", pgcLogPref, err)
	}

	if len(csvFiles) != 1 {
		t.Fatalf("%v: expected 1 csv file, not %v", pgcLogPref, len(csvFiles))
	}
}

// Test_ListAllBills checks if ListAllBills returns correct bills from db based on test data
func Test_ListAllBills(t *testing.T) {
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

	bills, err := pgcln.ListAllBills()
	if err != nil {
		t.Fatalf("%v: list all bills err: %v", pgcLogPref, err)
	}

	if len(bills) != 1 {
		t.Fatalf("%v: expected 1 bill, not %v", pgcLogPref, len(bills))
	}

	if strings.Compare(bills[0].LineItem, "test_service") != 0 {
		t.Fatalf("%v: bill's line item doesn't match 'test_service'", pgcLogPref)
	}
}

// Test_ListBillsByTime checks if ListBillsByTime() returns bills from non-empty db
func Test_ListBillsByTime(t *testing.T) {
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

	bills, err := pgcln.ListBillsByTime(start, end)
	if err != nil {
		t.Fatalf("%v: list bills err: %v", pgcLogPref, err)
	}

	if len(bills) == 0 {
		t.Fatalf("%v: expected non-empty list, but no bills were listed", pgcLogPref)
	}

	if strings.Compare(bills[0].LineItem, "test_service") != 0 {
		t.Fatalf("%v: bill's line item doesn't match 'test_service'", pgcLogPref)
	}
}

// Test_AddBill_removeLastBill tests adding bill into db, listing it by service and removing it
func Test_AddBill_removeLastBill(t *testing.T) {
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

	if err := pgcln.AddBill(testBill); err != nil {
		t.Fatalf("%v: add bill err: %v", pgcLogPref, err)
	}

	bills, err := pgcln.ListBillsByService("estIte")
	if err != nil {
		t.Fatalf("%v: list bills err: %v", pgcLogPref, err)
	}

	if len(bills) != 1 {
		t.Fatalf("%v: expected 1 bill, not %v", pgcLogPref, len(bills))
	}

	if strings.Compare(bills[0].LineItem, "testItem") != 0 || strings.Compare(bills[0].ProjectID, "testProject") != 0 {
		t.Fatalf("%v: bill's line item doesn't match the test bill", pgcLogPref)
	}

	if err := pgcln.removeLastBill(); err != nil {
		t.Fatalf("%v: remove last bill err: %v", pgcLogPref, err)
	}

	bills, err = pgcln.ListBillsByService("testItem")
	if err != nil {
		t.Fatalf("%v: list bills err: %v", pgcLogPref, err)
	}

	if len(bills) != 0 {
		t.Fatalf("%v: expected 0 bills, not %v", pgcLogPref, len(bills))
	}
}

// Test_ListBillsByProject checks if ListBillsByProject() returns correct bill from db by project
func Test_ListBillsByProject(t *testing.T) {
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

	bill, err := pgcln.ListBillsByProject("test_project")
	if err != nil {
		t.Fatalf("%v: list bills err: %v", pgcLogPref, err)
	}

	if strings.Compare(bill[0].LineItem, "test_project") != 0 {
		t.Fatalf("%v: bill's line item doesn't match 'test_service'", pgcLogPref)
	}
}

// Test_GetLastBill checks if GetLastBill() returns the last bill based on test data from db
func Test_GetLastBill(t *testing.T) {
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

	bill, err := pgcln.GetLastBill()
	if err != nil {
		t.Fatalf("%v: list bills err: %v", pgcLogPref, err)
	}

	if strings.Compare(bill.ProjectID, "test_project") != 0 {
		t.Fatalf("%v: bill's project id doesn't match 'test_project'", pgcLogPref)
	}
}
