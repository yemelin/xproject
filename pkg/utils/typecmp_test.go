package utils

import (
	"testing"
)

func TestMapsEquality(t *testing.T) {
	testcases := []struct {
		name   string
		in1    map[string]([]float64)
		in2    map[string]([]float64)
		result bool
	}{
		{
			name: "positive test",
			in1: map[string]([]float64){
				"a": []float64{0.0, 0.5, 1.0, 1.5},
				"b": []float64{1.0, 2.5, 1.3, 4.5},
			},
			in2: map[string]([]float64){
				"a": []float64{0.0, 0.5, 1.0, 1.5},
				"b": []float64{1.0, 2.5, 1.3, 4.5},
			},
			result: true,
		},
		{
			name: "negatitive test",
			in1: map[string]([]float64){
				"a": []float64{0.1, 0.5, 1.0, 1.5},
				"b": []float64{1.0, 2.8, 1.6, 4.5},
			},
			in2: map[string]([]float64){
				"a": []float64{0.1, 2.5, 1.3, 1.4},
				"b": []float64{1.0, 2.5, 1.3, 4.5},
			},
			result: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.result != IsMapsEqual(testcase.in1, testcase.in2) {
				t.Fatal("Slices arn't the same: ", testcase.in1, "!=", testcase.in2)
			}
		})
	}
}
