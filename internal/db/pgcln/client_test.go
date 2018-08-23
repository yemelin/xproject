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

	accounts, err := pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	prevLen := len(accounts)

	testAccount := GcpAccount{
		ID:             1,
		GcpAccountInfo: "testInfo",
	}

	if err := pgcln.AddAccount(testAccount); err != nil {
		t.Fatalf("%v: add account err: %v", pgcLogPref, err)
	}

	accounts, err = pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	if len(accounts)-prevLen != 1 {
		t.Fatalf("%v: expected 1 new account, not %v", pgcLogPref, len(accounts)-prevLen)
	}

	if strings.Compare(accounts[len(accounts)-1].GcpAccountInfo, "testInfo") != 0 {
		t.Fatalf("%v: account's info doesn't match the test account", pgcLogPref)
	}

	if err := pgcln.removeLastAccount(); err != nil {
		t.Fatalf("%v: remove last account err: %v", pgcLogPref, err)
	}

	accounts, err = pgcln.ListAccounts()
	if err != nil {
		t.Fatalf("%v: list accounts err: %v", pgcLogPref, err)
	}

	if len(accounts) != prevLen {
		if prevLen != 1 {
			t.Fatalf("%v: expected %v accounts, not %v", pgcLogPref, prevLen, len(accounts))
		} else {
			t.Fatalf("%v: expected %v account, not %v", pgcLogPref, prevLen, len(accounts))
		}
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

	csvFiles, err := pgcln.ListCsvFiles()
	if err != nil {
		t.Fatalf("%v: list csv file err: %v", pgcLogPref, err)
	}

	prevLen := len(csvFiles)

	testAccount := GcpAccount{
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

	testCsvFile := GcpCsvFile{
		ID:          1,
		Name:        "testName",
		Bucket:      "testBucket",
		TimeCreated: time.Date(2078, 1, 1, 0, 0, 0, 0, time.Local),
		AccountID:   accounts[len(accounts)-1].ID,
	}

	if err := pgcln.AddCsvFile(testCsvFile); err != nil {
		t.Fatalf("%v: add csv file err: %v", pgcLogPref, err)
	}

	csvFiles, err = pgcln.ListCsvFiles()
	if err != nil {
		t.Fatalf("%v: list csv file err: %v", pgcLogPref, err)
	}

	if len(csvFiles)-prevLen != 1 {
		t.Fatalf("%v: expected 1 new csv file, not %v", pgcLogPref, len(csvFiles)-prevLen)
	}

	if strings.Compare(csvFiles[len(csvFiles)-1].Name, "testName") != 0 {
		t.Fatalf("%v: csv file's name doesn't match the test csv file", pgcLogPref)
	}

	if err := pgcln.removeLastCsvFile(); err != nil {
		t.Fatalf("%v: remove last csv file err: %v", pgcLogPref, err)
	}

	csvFiles, err = pgcln.ListCsvFiles()
	if err != nil {
		t.Fatalf("%v: list csv files err: %v", pgcLogPref, err)
	}

	if len(csvFiles) != prevLen {
		if prevLen != 1 {
			t.Fatalf("%v: expected %v csv files, not %v", pgcLogPref, prevLen, len(csvFiles))
		} else {
			t.Fatalf("%v: expected %v csv file, not %v", pgcLogPref, prevLen, len(csvFiles))
		}
	}

	if err := pgcln.removeLastAccount(); err != nil {
		t.Fatalf("%v: remove last account err: %v", pgcLogPref, err)
	}
}

// Test_Bill tests all functions that are related to adding, listing and removing service bills
func Test_Bill(t *testing.T) {
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

	prevLen := len(bills)

	testAccount := GcpAccount{
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

	testCsvFile := GcpCsvFile{
		ID:          1,
		Name:        "testName",
		Bucket:      "testBucket",
		TimeCreated: time.Date(2078, 1, 1, 0, 0, 0, 0, time.Local),
		AccountID:   accounts[len(accounts)-1].ID,
	}

	if err := pgcln.AddCsvFile(testCsvFile); err != nil {
		t.Fatalf("%v: add csv file err: %v", pgcLogPref, err)
	}

	csvFiles, err := pgcln.ListCsvFiles()
	if err != nil {
		t.Fatalf("%v: list csv file err: %v", pgcLogPref, err)
	}

	testBill1 := ServiceBill{
		ID:           1,
		LineItem:     "testItem1",
		StartTime:    time.Date(2077, 1, 1, 0, 0, 0, 0, time.Local),
		EndTime:      time.Date(2077, 1, 1, 1, 0, 0, 0, time.Local),
		Cost:         123.456,
		Currency:     "testCurrency1",
		ProjectID:    "testProject1",
		Description:  "testDescription1",
		GcpCsvFileID: csvFiles[len(csvFiles)-1].ID,
	}

	testBill2 := ServiceBill{
		ID:           2,
		LineItem:     "testItem2",
		StartTime:    time.Date(2078, 1, 1, 0, 0, 0, 0, time.Local),
		EndTime:      time.Date(2078, 1, 1, 1, 0, 0, 0, time.Local),
		Cost:         456.789,
		Currency:     "testCurrency2",
		ProjectID:    "testProject2",
		Description:  "testDescription2",
		GcpCsvFileID: csvFiles[len(csvFiles)-1].ID,
	}

	if err := pgcln.AddBill(testBill1); err != nil {
		t.Fatalf("%v: add bill err: %v", pgcLogPref, err)
	}

	if err := pgcln.AddBill(testBill2); err != nil {
		t.Fatalf("%v: add bill err: %v", pgcLogPref, err)
	}

	lastBill, err := pgcln.GetLastBill()
	if err != nil {
		t.Fatalf("%v: get last bill err: %v", pgcLogPref, err)
	}

	if strings.Compare(testBill2.LineItem, lastBill.LineItem) != 0 || strings.Compare(testBill2.ProjectID, lastBill.ProjectID) != 0 {
		t.Fatalf("%v: last bill doesn't match the test bill", pgcLogPref)
	}

	bills, err = pgcln.ListBillsByService("testItem")
	if err != nil {
		t.Fatalf("%v: list bills by service err: %v", pgcLogPref, err)
	}

	if len(bills) != 2 {
		t.Fatalf("%v: expected 2 bills, not %v", pgcLogPref, len(bills))
	}

	bills, err = pgcln.ListBillsByTime(time.Date(2077, 12, 31, 0, 0, 0, 0, time.Local), time.Date(2078, 1, 15, 0, 0, 0, 0, time.Local))
	if err != nil {
		t.Fatalf("%v: list bills by time err: %v", pgcLogPref, err)
	}

	if len(bills) != 1 {
		t.Fatalf("%v: expected 1 bill, not %v", pgcLogPref, len(bills))
	}

	if strings.Compare(testBill2.Description, bills[0].Description) != 0 {
		t.Fatalf("%v: bill's description doesn't match the test bill", pgcLogPref)
	}

	bills, err = pgcln.ListBillsByProject("testProject1")
	if err != nil {
		t.Fatalf("%v: list bills by project err: %v", pgcLogPref, err)
	}

	if len(bills) != 1 {
		t.Fatalf("%v: expected 1 bill, not %v", pgcLogPref, len(bills))
	}

	if strings.Compare(testBill1.Description, bills[0].Description) != 0 {
		t.Fatalf("%v: bill's description doesn't match the test bill", pgcLogPref)
	}

	for i := 0; i < 2; i++ {
		if err := pgcln.removeLastBill(); err != nil {
			t.Fatalf("%v: remove last bill err: %v", pgcLogPref, err)
		}
	}

	bills, err = pgcln.ListAllBills()
	if err != nil {
		t.Fatalf("%v: list all bills err: %v", pgcLogPref, err)
	}

	if len(bills) != prevLen {
		if prevLen != 1 {
			t.Fatalf("%v: expected %v bills, not %v", pgcLogPref, prevLen, len(bills))
		} else {
			t.Fatalf("%v: expected %v bill, not %v", pgcLogPref, prevLen, len(bills))
		}
	}

	if err := pgcln.removeLastCsvFile(); err != nil {
		t.Fatalf("%v: remove last csv file err: %v", pgcLogPref, err)
	}

	if err := pgcln.removeLastAccount(); err != nil {
		t.Fatalf("%v: remove last account err: %v", pgcLogPref, err)
	}
}
