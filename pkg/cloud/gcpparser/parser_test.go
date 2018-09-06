// Package gcpparser uses for parse raw GCP billing csv file
package gcpparser

import (
	"encoding/csv"
	"os"
	"testing"
)

// Test_Parse tests Parse function
func Test_Parse(t *testing.T) {

	// opening and reading test data
	f, err := os.Open("./testdata/test1.csv")
	if err != nil {
		t.Errorf("Can not open test1.csv: %v", err)
	}
	records, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	// parse test data
	res, err := Parse(records)
	if err != nil {
		t.Errorf("err != nil: %v", err)
	}

	// comparing record with data from "./testdata/test1.csv"
	if len(res) == 0 {
		t.Errorf("Got: Res len < 0, Exp: len > 0")
		t.FailNow()
	}
	exp := "com.google.cloud/services/cloud-storage/StorageMultiRegionalUsGbsec"
	if res[0].LineItem != exp {
		t.Errorf("parse:\nExp: %v, Got: %v", exp, res[0].LineItem)
	}

	// now we are opening and reading test csv file contains only header
	f, err = os.Open("./testdata/test_only_header.csv")
	if err != nil {
		t.Errorf("Can not open test_only_header.csv: %v", err)
	}
	records, err = csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	// parse test data
	res, err = Parse(records)
	if err != nil {
		t.Errorf("err != nil: %v", err)
	}

	// compare results
	if len(res) != 0 {
		t.Errorf("Got: Res len > 0, Exp: len < 0")
		t.FailNow()
	}

	// here test csv file contains wrong header
	f, err = os.Open("./testdata/test_wrong_header.csv")
	if err != nil {
		t.Errorf("Can not open test_wrong_header.csv: %v", err)
	}
	records, err = csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		t.Errorf("Can not parse file from bucket: %v", err)
	}

	// parse test data
	_, err = Parse(records)
	if err == nil {
		t.Error("Parsed file with wrong header")
	}
	// test expected nil argument
	_, err = Parse(nil)
	if err == nil {
		t.Error("Parse nil argument")
	}

	// parse file with nil headers
	_, err = Parse([][]string{nil, nil, nil, nil, nil, nil, nil, nil})
	if err == nil {
		t.Errorf("Error while parsing [][]string{nil, nil, nil, nil, nil, nil, nil, nil}, expected error")
	}
}
