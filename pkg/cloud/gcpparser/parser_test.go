// Package gcpparser uses for parse raw GCP billing csv file
package gcpparser

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"
)

func Test_Parse(t *testing.T) {
	f, _ := os.Open("./testdata/test1.csv")
	r := bufio.NewReader(f)
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	const minRecNum = 1
	if len(records) < minRecNum {
		t.Errorf("records number < 1")
	}
	// skip row with headers
	res, err := Parse(records[minRecNum:])
	if err != nil {
		t.Errorf("err != nil: %v", err)
	}
	// compare record with data from "./testdata/test1.csv"
	exp := "com.google.cloud/services/cloud-storage/StorageMultiRegionalUsGbsec"
	if res[0].Item != exp {
		t.Errorf("parse:\nExp: %v, Got: %v", exp, res[0].Item)
	}
}
