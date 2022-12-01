package utils

import (
	"strings"
)

// TagMatch 標籤符合
func TagMatch(lhs, rhs string) bool {
	for _, itor := range lhs {
		if strings.ContainsRune(rhs, itor) {
			return true
		} // if
	} // for

	return false
}
