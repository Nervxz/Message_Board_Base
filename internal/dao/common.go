package dao

import (
	"strconv"
	"strings"
)

// genSQLParams takes the number of records and the number of arguments per record
// generates the SQL parameters for the query
// numRecords: The number of records for which to generate SQL
// sb: A pointer to a strings.Builder used to build the SQL string
// numArgs: The number of arguments per record
func genSQLParams(numRecords int, sb *strings.Builder, numArgs int) {
	for i := 0; i < numRecords; i++ {
		if i != 0 {
			// Add a comma to separate the records
			sb.WriteString(",")
		}

		id := i*numArgs + 1
		sb.WriteString("($")
		sb.WriteString(strconv.Itoa(id))

		for j := 1; j < numArgs; j++ {
			sb.WriteString(",$")
			sb.WriteString(strconv.Itoa(j + id))
		}

		sb.WriteString(")")
	}
}
