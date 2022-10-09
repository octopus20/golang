package mystring

import (
	"sort"
)

// Sort

func SortString(stringarray []string) []string {
	{
		sort.Strings(stringarray)
		return stringarray
	}
}
