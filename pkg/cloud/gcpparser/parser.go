package gcpparser

// This package need for parsing gcp billing csv data
// Also this package store

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pavlov-tony/xproject/pkg/cloud/gcptypes"
)

// necessary columns
const (
	lineItem    = "Line Item"
	startTime   = "Start Time"
	endTime     = "End Time"
	cost        = "Cost"
	currency    = "Currency"
	projectID   = "Project ID"
	description = "Description"
)

// array of necessary columns
var necesCols = [...]string{lineItem, startTime, endTime, cost, currency, projectID, description}

// Parse need for parse full GCP csv billing file
func Parse(data [][]string) (res gcptypes.ServicesBills, err error) {

	// check if input data == nil
	if data == nil {
		return nil, fmt.Errorf("parse: data == nil")
	}
	// check if input data has not columns
	if len(data) == 0 {
		return nil, fmt.Errorf("parse: data len == 0")
	}

	// mark necessary headers (targets)
	tg, err := markCols(data[0])
	if err != nil {
		return nil, fmt.Errorf("parse: error in markCols: %v", err)
	}

	// skip headers, handling content line by line
	data = data[1:]
	for _, l := range data {
		sb, err := parseLine(l, tg)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		res = append(res, sb)
	}

	return res, nil
}

// parseLine need for parse row from GCP csv billing file
func parseLine(line []string, tg targCols) (*gcptypes.ServiceBill, error) {

	// parse startTime and endTime
	st, err := time.Parse(time.RFC3339, line[tg[startTime]])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse StartTime, %v", err)
	}
	et, err := time.Parse(time.RFC3339, line[tg[endTime]])
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse EndTime, %v", err)
	}
	// parse cost
	cst, err := strconv.ParseFloat(line[tg[cost]], 64)
	if err != nil {
		return nil, fmt.Errorf("parse line: can not parse Cost, %v", err)
	}

	// create new ServiceBill
	sb := gcptypes.ServiceBill{
		LineItem:    line[tg[lineItem]],
		StartTime:   st,
		EndTime:     et,
		Cost:        cst,
		Currency:    line[tg[currency]],
		ProjectID:   line[tg[projectID]],
		Description: line[tg[description]],
	}

	return &sb, nil
}

// markCols takes headers of billing file and returns numbers of necessary columns
func markCols(headers []string) (targCols, error) {

	res := make(targCols)

	// match headers from file with necessary columns
	for i, h := range headers {
		for _, v := range necesCols {
			if h == v {
				res[v] = i
			}
		}
	}

	// if all necessary columns have not matched
	if len(res) != len(necesCols) {
		return nil, fmt.Errorf("can not mark all columns, marked %v", res)
	}

	return res, nil
}
