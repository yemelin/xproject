package reportrow

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ReportRow represents the record exctracted from CSV report and ready for inserting to db
// This is the prototype of the actual DbRecord. Its design should be discussed.
type ReportRow struct {
	// Unique operation id
	IdentityLineItemId string

	// Time intervals
	TimeIntervalStart time.Time
	TimeIntervalEnd   time.Time

	BillPayerAccountId uint64

	// BillingPeriod refers to the current month
	BillBillingPeriodStartDate time.Time
	BillBillingPeriodEndDate   time.Time

	// Possible values: "Usage", "Tax". "Tax" is the only row in the CSV
	LineItemLineItemType string

	// ProductCode is almost the same as ProductName, but shorter
	// E.g. "AmazonEC2" vs "Amazon Elastic Cloud Computing"
	// Possible values: "AWSCostExplorer", "AmazonCloudWatch", "AmazonEC2", ...
	LineItemProductCode string
	LineItemUsageAmount float64

	// Possible values: USD, ...
	LineItemCurrencyCode string

	LineItemUnblendedRate float64
	LineItemUnblendedCost float64
	LineItemBlendedRate   float64
	LineItemBlendedCost   float64

	// Possible values: <blank>, "eu-central-1", "us-east-1", ...
	Productregion string

	// Product SKU. Unique product id
	Productsku string

	// Public on demand cost and rate
	PricingpublicOnDemandCost float64
	PricingpublicOnDemandRate float64

	// Possible values: <blank>, "OnDemand", ...
	Pricingterm string
	// Possible values: <blank>, "GB", "Request", ...
	Pricingunit string
}

// FromStrings parses a ReportRow from a slice of strings
func FromStrings(row []string) (*ReportRow, error) {
	IdentityLineItemId := row[0]

	timeInterval := row[1]
	intervals := strings.Split(timeInterval, "/")
	start := intervals[0]
	end := intervals[1]
	TimeIntervalStart, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return nil, fmt.Errorf("can't parse TimeIntervalStart: %v", err)
	}
	TimeIntervalEnd, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return nil, fmt.Errorf("can't parse TimeIntervalEnd: %v", err)
	}

	BillPayerAccountId, err := strconv.ParseUint(row[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as uint64 BillPayerAccountId: %v", err)
	}

	BillBillingPeriodStartDate, err := time.Parse(time.RFC3339, row[3])
	if err != nil {
		return nil, fmt.Errorf("can't parse BillBillingPeriodStartDate: %v", err)
	}
	BillBillingPeriodEndDate, err := time.Parse(time.RFC3339, row[4])
	if err != nil {
		return nil, fmt.Errorf("can't parse BillBillingPeriodEndDate: %v", err)
	}

	LineItemLineItemType := row[5]
	LineItemProductCode := row[6]
	LineItemUsageAmount, err := strconv.ParseFloat(row[7], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as float64 LineItemUsageAmount: %v", err)
	}

	LineItemCurrencyCode := row[8]

	LineItemUnblendedRate, err := strconv.ParseFloat(row[9], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as float64 LineItemUsageAmount: %v", err)
	}

	LineItemUnblendedCost, err := strconv.ParseFloat(row[10], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as float64 LineItemUnblendedCost: %v", err)
	}

	LineItemBlendedRate, err := strconv.ParseFloat(row[11], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as float64 LineItemBlendedRate: %v", err)
	}

	LineItemBlendedCost, err := strconv.ParseFloat(row[12], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as float64 LineItemBlendedCost: %v", err)
	}

	Productregion := row[13]
	Productsku := row[14]

	PricingpublicOnDemandCost, err := strconv.ParseFloat(row[15], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as float64 PricingpublicOnDemandCost: %v", err)
	}
	PricingpublicOnDemandRate, err := strconv.ParseFloat(row[16], 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse as float64 PricingpublicOnDemandRate: %v", err)
	}

	Pricingterm := row[17]
	Pricingunit := row[18]

	return &ReportRow{
		IdentityLineItemId:         IdentityLineItemId,
		TimeIntervalStart:          TimeIntervalStart,
		TimeIntervalEnd:            TimeIntervalEnd,
		BillPayerAccountId:         BillPayerAccountId,
		BillBillingPeriodStartDate: BillBillingPeriodStartDate,
		BillBillingPeriodEndDate:   BillBillingPeriodEndDate,
		LineItemLineItemType:       LineItemLineItemType,
		LineItemProductCode:        LineItemProductCode,
		LineItemUsageAmount:        LineItemUsageAmount,
		LineItemCurrencyCode:       LineItemCurrencyCode,
		LineItemUnblendedRate:      LineItemUnblendedRate,
		LineItemUnblendedCost:      LineItemUnblendedCost,
		LineItemBlendedRate:        LineItemBlendedRate,
		LineItemBlendedCost:        LineItemBlendedCost,
		Productregion:              Productregion,
		Productsku:                 Productsku,
		PricingpublicOnDemandCost:  PricingpublicOnDemandCost,
		PricingpublicOnDemandRate:  PricingpublicOnDemandRate,
		Pricingterm:                Pricingterm,
		Pricingunit:                Pricingunit,
	}, nil

}
