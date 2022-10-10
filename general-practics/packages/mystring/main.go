package mystring

import (
	"fmt"
	"sort"
)

// Sort

func SortString(stringarray []string) []string {
	{
		sort.Strings(stringarray)
		return stringarray
	}
}

func Occurrences(st string) string {
	charc := map[rune]int{}

	chars := []rune{}
	for char := range charc {
		chars = append(chars, char)
	}
	for _, char := range chars {
		fmt.Printf("%c %d\n", char, charc[char])
	}
	// return fmt.Sprintf("%c %d\n", char, charc[char])
}
