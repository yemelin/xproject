package gcpparser

import (
	"errors"
	"strconv"
	"time"
)

func Parse(data [][]string) (res ServicesBills, err error) {
	for _, l := range data {
        sb, err := 
	}

	return res, nil
}

func parseLine(line []string) (*ServiceBill, error) {
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
