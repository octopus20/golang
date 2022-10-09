package mystring

import (
	"fmt"
	"sort"
)

// Sort

func SortString() {

	stringarray := []string{"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "A", "W", "E", "I", "O", "U", "P", "Q", "R", "S", "T", "V", "X", "Z"}
	{
		sort.Strings(stringarray)
		fmt.Println("Strings:", stringarray)
	}
}
