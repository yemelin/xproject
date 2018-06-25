package csvreader

import (
	"io"
	"strings"
	"testing"

	"github.com/pavlov-tony/xproject/pkg/csvparse/csvparseutils"
)

func TestReadRow(t *testing.T) {
	testcases := []struct {
		name    string
		in      string
		columns []string
		out     [][]string
	}{
		{
			name: "Filter by first_name and username",
			in: `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`,
			columns: []string{"first_name", "username"},
			out: [][]string{
				{"Rob", "rob"},
				{"Ken", "ken"},
				{"Robert", "gri"},
			},
		},
		{
			name: "Filter by a, c, and d",
			in: `a,b,c,d,e
1,2,3,4,5
6,7,8,9,0
1,1,1,1,1
`,
			columns: []string{"a", "c", "d"},
			out: [][]string{
				{"1", "3", "4"},
				{"6", "8", "9"},
				{"1", "1", "1"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			r := New(strings.NewReader(testcase.in), testcase.columns)

			for _, row := range testcase.out {
				s, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("Error while reading from csv: %v", err)
				}
				if !csvparseutils.AreStringSlicesEqual(s, row) {
					t.Fatalf("Rows are not equal: %v != %v, (%v, %v)", s, row, len(s), len(row))
				}
			}
		})
	}
}

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
