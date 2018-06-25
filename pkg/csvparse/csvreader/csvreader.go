package csvreader

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CsvReader struct {
	filter    []string
	indices   []int
	csvReader *csv.Reader
}

// New returns a new CsvReader
// Note that it skips the header for purpose
func New(r io.Reader, filterHeaders []string) *CsvReader {
	return &CsvReader{
		filter:    filterHeaders,
		indices:   nil,
		csvReader: csv.NewReader(r),
	}
}

// ColumnIndexByName returns column index from provided headers by its name.
// If there are several headers with the same names, returns the first.
func ColumnIndexByName(header []string, name string) (int, error) {
	for i, colName := range header {
		if colName == name {
			return i, nil
		}
	}

	return 0, fmt.Errorf("header with name \"%v\" was not found", name)
}

// Read is actually reads a row from CSV as a slice of strings
func (r *CsvReader) Read() ([]string, error) {
	if r.indices == nil {
		r.indices = make([]int, len(r.filter))

		header, err := r.csvReader.Read()
		if err == io.EOF {
			return nil, err
		}
		if err != nil {
			return nil, fmt.Errorf("can't read header: %v", err)
		}

		// find indices from filter in header
		for i, columnName := range r.filter {
			idx, err := ColumnIndexByName(header, columnName)
			if err != nil {
				return nil, fmt.Errorf("can't get column index by name \"%v\": %v", columnName, err)
			}
			r.indices[i] = idx
		}
	}

	record, err := r.csvReader.Read()
	// fmt.Println("rec len:", len(record))
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("can't read row: %v", err)
	}

	// fmt.Println("r.filt len:", len(r.filter))
	filteredRecord := make([]string, len(r.filter))
	// fmt.Println("filt rec len:", len(filteredRecord))

	for i, idx := range r.indices {
		filteredRecord[i] = record[idx]
	}

	return filteredRecord, nil
}
