package csvutils

import (
	"testing"
)

func TestColumnIndexByName(t *testing.T) {
	testcases := []struct {
		name   string
		in     string
		header []string
		out    int
	}{
		{
			name:   "Get the index of header 'a'",
			in:     "a",
			header: []string{"a", "b", "c"},
			out:    0,
		},
		{
			name:   "Get the index of the first header 'smth/with/slashes'",
			in:     "smth/with/slashes",
			header: []string{"a", "b", "c", "smth/with/slashes", "smth/with/slashes"},
			out:    3,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			r, err := ColumnIndexByName(testcase.header, testcase.in)
			if err != nil {
				t.Fatalf("Error while getting the header index by name '%v': %v", testcase.in, err)
			}

			if r != testcase.out {
				t.Fatalf("Indices are not equal: %v != %v", r, testcase.out)
			}
		})
	}
}
