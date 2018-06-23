package model

import "time"

// Important fields for prediction from gcp billing csv
type ServiceBill struct {
	Description string // NOTE: Now field 17, may be should change
	StartTime   time.Time
	EndTime     time.Time
	Cost        float64
	Currency    string
}

// Slice of Services Bill
type ServicesBill []ServiceBill

// Get Service bill from Services Bill slice
func (ssb ServicesBill) getBillById() ServiceBill {
	return ssb[1]
}
