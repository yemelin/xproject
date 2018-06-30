package gcpparser

import (
	"errors"
	"strconv"
	"time"
)

// Parse full GCP csv billing file
func Parse(data [][]string) (res ServicesBills, err error) {
	for _, l := range data {
		sb, err := parseLine(l)
		if err != nil {
			return nil, errors.New("Parse: " + err.Error())
		}
		res = append(res, sb)
	}

	return res, nil
}

// Parse row from GCP csv billing file
func parseLine(line []string) (*ServiceBill, error) {
	// TODO: fix magic number
	if len(line) < 20 {
		return nil, errors.New("parseLine: line length < 20")
	}

	st, err := time.Parse(time.RFC3339, line[ColStartTime])
	if err != nil {
		return nil, errors.New("parseLine: can not parse StartTime")
	}
	et, err := time.Parse(time.RFC3339, line[ColEndTime])
	if err != nil {
		return nil, errors.New("parseLine: can not parse EndTime")
	}
	cst, err := strconv.ParseFloat(line[ColCost], 64)
	if err != nil {
		return nil, errors.New("parseLine: can not parse Cost")
	}

	return &ServiceBill{
		Item:    line[ColLineItem],
		Started: st,
		Ended:   et,
		Cost:    cst,
		Curr:    line[ColCurrency],
		ProjID:  line[ColProjectID],
		Descr:   line[ColDescription],
	}, nil
}
