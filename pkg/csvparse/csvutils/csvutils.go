package csvutils

import (
	"fmt"
)

// ColumnIndexByName returns a column index of the header with a
// provided name. If there are several headers with the same names,
// returns the first.
func ColumnIndexByName(header []string, name string) (int, error) {
	for i, colName := range header {
		if colName == name {
			return i, nil
		}
	}

	return 0, fmt.Errorf("column with name \"%v\" was not found", name)
}
