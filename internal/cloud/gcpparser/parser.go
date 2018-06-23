package gcpparser

import (
	"time"
)

type CsvRows [][]string

type ServiceBill struct {
	Description string // NOTE: Now field 17, may be should change
	StartTime   time.Time
	EndTime     time.Time
	Cost        float64
	Currency    string
}

type ServicesBill [][]string

func parseServicesBill(CsvObject [][]string) {

}
