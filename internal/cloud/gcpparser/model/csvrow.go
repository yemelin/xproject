package model

type CsvObject struct {
	records [][]string
}

func (o CsvObject) Row(i int) []string {
	return o.records[i]
}
