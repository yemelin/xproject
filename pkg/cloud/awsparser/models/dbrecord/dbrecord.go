package dbrecord

import (
	"time"
)

// DbRecord represents the record exctracted from CSV report and ready for inserting to db
// This is the prototype of the actual DbRecord. Its design should be discussed.
type DbRecord struct {
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
