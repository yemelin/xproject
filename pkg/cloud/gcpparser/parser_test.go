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
	defer f.Close()
	r := bufio.NewReader(f)
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	const minRecNum = 1
	if len(records) < minRecNum {
		t.Errorf("records number < 1")
	}

	res, err := Parse(records)
	if err != nil {
		t.Errorf("err != nil: %v", err)
	}
	// compare record with data from "./testdata/test1.csv"
	exp := "com.google.cloud/services/cloud-storage/StorageMultiRegionalUsGbsec"
	if res[0].LineItem != exp {
		t.Errorf("parse:\nExp: %v, Got: %v", exp, res[0].LineItem)
	}
}

func Test_markCols(t *testing.T) {
	f, _ := os.Open("./testdata/test1.csv")
	defer f.Close()
	r := bufio.NewReader(f)
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		t.Errorf("can not parse file from bucket: %v", err)
	}

	tg, err := markCols(records[0])
	if err != nil {
		t.Errorf("can not mark main headers: %v", err)
	}
	// main columns from test csv
	exp := targCols{1, 2, 3, 14, 15, 17, 20}
	if *tg != exp {
		t.Errorf("markCols: Exp: %v, Got: %v", exp, *tg)
	}
}
