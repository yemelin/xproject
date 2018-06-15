package awsparser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pavlov-tony/xproject/pkg/cloud/awsparser/models/dbrecord"
)

// RawCsvToDbRecords converts RawCsv struct to the slice of DbRecord
func RawCsvToDbRecords(csv *RawCsv) ([]*dbrecord.DbRecord, error) {
	filteredCsv, err := csv.FilterByNames([]string{
		"identity/LineItemId",
		"identity/TimeInterval", // should br splitted to start and end
		"bill/PayerAccountId",
		"bill/BillingPeriodStartDate",
		"bill/BillingPeriodEndDate",
		"lineItem/LineItemType",
		"lineItem/ProductCode",
		"lineItem/UsageAmount",
		"lineItem/CurrencyCode",
		"lineItem/UnblendedRate",
		"lineItem/UnblendedCost",
		"lineItem/BlendedRate",
		"lineItem/BlendedCost",
		"product/region",
		"product/sku",
		"pricing/publicOnDemandCost",
		"pricing/publicOnDemandRate",
		"pricing/term",
		"pricing/unit",
	})

	if err != nil {
		return nil, fmt.Errorf("can't filter csv by required columns: %v", err)
	}

	result := make([]*dbrecord.DbRecord, len(filteredCsv.Rows()))

	for _, filtered := range filteredCsv.Rows()[1:] {

		IdentityLineItemId := filtered[0]

		timeInterval := filtered[1]
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

		BillPayerAccountId, err := strconv.ParseUint(filtered[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as uint64 BillPayerAccountId: %v", err)
		}

		BillBillingPeriodStartDate, err := time.Parse(time.RFC3339, filtered[3])
		if err != nil {
			return nil, fmt.Errorf("can't parse BillBillingPeriodStartDate: %v", err)
		}
		BillBillingPeriodEndDate, err := time.Parse(time.RFC3339, filtered[4])
		if err != nil {
			return nil, fmt.Errorf("can't parse BillBillingPeriodEndDate: %v", err)
		}

		LineItemLineItemType := filtered[5]
		LineItemProductCode := filtered[6]
		LineItemUsageAmount, err := strconv.ParseFloat(filtered[7], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as float64 LineItemUsageAmount: %v", err)
		}

		LineItemCurrencyCode := filtered[8]

		LineItemUnblendedRate, err := strconv.ParseFloat(filtered[9], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as float64 LineItemUsageAmount: %v", err)
		}

		LineItemUnblendedCost, err := strconv.ParseFloat(filtered[10], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as float64 LineItemUnblendedCost: %v", err)
		}

		LineItemBlendedRate, err := strconv.ParseFloat(filtered[11], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as float64 LineItemBlendedRate: %v", err)
		}

		LineItemBlendedCost, err := strconv.ParseFloat(filtered[12], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as float64 LineItemBlendedCost: %v", err)
		}

		Productregion := filtered[13]
		Productsku := filtered[14]

		PricingpublicOnDemandCost, err := strconv.ParseFloat(filtered[15], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as float64 PricingpublicOnDemandCost: %v", err)
		}
		PricingpublicOnDemandRate, err := strconv.ParseFloat(filtered[16], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse as float64 PricingpublicOnDemandRate: %v", err)
		}

		Pricingterm := filtered[17]
		Pricingunit := filtered[18]

		dbrec := &dbrecord.DbRecord{
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
		}

		result = append(result, dbrec)
	}

	return result, nil
}
