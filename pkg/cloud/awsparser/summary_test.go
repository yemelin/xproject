package awsparser

import (
	"testing"

	"github.com/pavlov-tony/xproject/pkg/cloud/awsparser/csvparseutils"
)

func TestGetSummary(t *testing.T) {
	testcases := []struct {
		name string
		in   *RawCsv
		cfg  *SummaryConfig
		out  *CsvSummary
	}{
		{
			name: "GroupBy 'a' and SumBy 'b'",
			in: &RawCsv{
				records: [][]string{
					{"a", "b", "c"},
					{"z", "2", "3"},
					{"z", "2", "3"},
					{"y", "5", "6"},
					{"y", "5", "6"},
				},
			},
			cfg: NewSummaryConfig("a", []string{"b"}),
			out: &CsvSummary{
				headers: []string{"a", "b"},
				inner: &map[string]([]float64){
					"z": []float64{4.0},
					"y": []float64{10.0},
				},
			},
		},
		{
			name: "GroupBy 'identity/TimeInterval' and SumBy 'lineItem/UnblendedCost'",
			in: &RawCsv{
				records: [][]string{
					{"identity/TimeInterval", "lineItem/UnblendedCost"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "1"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "2"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "3"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "4"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "10"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "20"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "30"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "40"},
				},
			},
			cfg: NewSummaryConfig("identity/TimeInterval", []string{"lineItem/UnblendedCost"}),
			out: &CsvSummary{
				headers: []string{"identity/TimeInterval", "lineItem/UnblendedCost"},
				inner: &map[string]([]float64){
					"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z": []float64{10},
					"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z": []float64{100},
				},
			},
		},
		{
			name: "GroupBy 'lineItem/ProductCode' and SumBy 'lineItem/UnblendedCost'",
			in: &RawCsv{
				records: [][]string{
					{"lineItem/ProductCode", "lineItem/UnblendedCost"},
					{"AmazonEC2", "1"},
					{"AmazonQuickSight", "2"},
					{"AmazonRDS", "3"},
					{"AmazonQuickSight", "4"},
					{"AmazonEC2", "10"},
					{"AmazonEC2", "20"},
					{"AmazonRDS", "30"},
					{"AmazonCloudWatch", "40"},
				},
			},
			cfg: NewSummaryConfig("lineItem/ProductCode", []string{"lineItem/UnblendedCost"}),
			out: &CsvSummary{
				headers: []string{"lineItem/ProductCode", "lineItem/UnblendedCost"},
				inner: &map[string]([]float64){
					"AmazonCloudWatch": []float64{40},
					"AmazonEC2":        []float64{31},
					"AmazonRDS":        []float64{33},
					"AmazonQuickSight": []float64{6},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			scfg := testcase.cfg
			summary, err := testcase.in.GetSummary(scfg)
			if err != nil {
				t.Fatal(err)
			}

			if !csvparseutils.AreStringSlicesEqual(summary.Headers(), testcase.out.Headers()) {
				t.Fatal("Header differs:", summary.Headers(), "!=", testcase.out.Headers())
			}
			if !csvparseutils.AreMapsEqual(*summary.Inner(), *testcase.out.Inner()) {
				t.Fatal("Summary differs:", *summary.Inner(), "!=", *testcase.out.Inner())
			}
		})
	}
}
