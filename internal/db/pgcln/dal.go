package pgcln

import "strconv"

// selectFrom returns query that selects all rows from table
func selectFrom(table string) string {
	return "SELECT * FROM " + table + " ORDER BY id ASC"
}

// selectByRange returns query that selects rows that are withing the specified range
func selectByRange(table, startCol, endCol string) string {
	return "SELECT * FROM " + table + " WHERE " + startCol + " >= $1 AND " + endCol + " <= $2 ORDER BY id ASC"
}

// selectByMatch returns query that selects rows that match the specified string
func selectByMatch(table, col string) string {
	return "SELECT * FROM " + table + " WHERE " + col + " LIKE $1 ORDER BY id ASC"
}

// selectLast returns query that selects row by largest value in col1, col2 and then col3
func selectLast(table, col1, col2, col3 string) string {
	return "SELECT * FROM " + table + " ORDER BY " + col1 + " DESC, " + col2 + " DESC, " + col3 + " DESC LIMIT 1"
}

// insertInto returns query that inserts a row into table with specified number of columns
func insertInto(table string, cols int) string {
	query := "INSERT INTO " + table + " VALUES(DEFAULT"

	for i := 0; i < cols; i++ {
		query += ", $" + strconv.Itoa(i+1)
	}

	query += ")"

	return query
}

// deleteFrom returns query that deletes row with maximum id (latest added row) from table
func deleteFrom(table string) string {
	return "DELETE FROM " + table + " WHERE id = (SELECT MAX(id) FROM " + table + ")"
}
