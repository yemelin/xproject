// Package for parsing Cost and Usage Reports manifest file
package manifest

import (
	"fmt"
)

// Column represents manifest column format
type Column struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

// Returns "category/name" for Column type
func (c Column) String() string {
	return fmt.Sprintf("%v/%v", c.Category, c.Name)
}

// BillingPeriod represents minifest billing period
type BillingPeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Manifest represents Cost and Usage reports manifest format
type Manifest struct {
	AssemblyId             string        `json:"assemblyId"`
	Account                string        `json:"account"`
	Columns                []Column      `json:"columns"`
	Charset                string        `json:"charset"`
	Compression            string        `json:"compression"`
	ContentType            string        `json:"contentType"`
	ReportId               string        `json:"reportId"`
	ReportName             string        `json:"reportName"`
	BillingPeriod          BillingPeriod `json:"billingPeriod"`
	Bucket                 string        `json:"bucket"`
	ReportKeys             []string      `json:"reportKeys"`
	AdditionalArtifactKeys []string      `json:"additionalArtifactKeys"`
}
