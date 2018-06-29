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
	records, _ := csv.NewReader(r).ReadAll()

	Parse(records)
}
