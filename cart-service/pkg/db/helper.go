package db

import (
	"fmt"
	"strings"
)

func idsToString(ids []int64, start int) string {
	placeholder := make([]string, len(ids))
	for i := range ids {
		placeholder[i] = fmt.Sprintf("$%d", i+start)
	}
	return strings.Join(placeholder, ",")
}
