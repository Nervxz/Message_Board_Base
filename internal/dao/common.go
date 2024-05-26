package dao

import (
	"strconv"
	"strings"
)

func genSQLParams(numRecords int, sb *strings.Builder, numArgs int) {
	for i := 0; i < numRecords; i++ {
		if i != 0 {
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
