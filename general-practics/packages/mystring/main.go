package mystring

import (
	"fmt"
	"sort"
	"unicode"
)

// Sort

func SortString(stringarray []string) []string {
	// COMMENT: what's the need for the {} ?
	{
		sort.Strings(stringarray)
		// COMMENT: the above sort.Strings(stringarray) does sort in place, so no need to return the slice again
		// COMMENT: Idealy, copy the original slice into new one, sort it and return it
		return stringarray
	}
}

// COMMENT: it would be better to return a map and leave it to the caller on how to represent the result
func Occurrences(st string) string {
	wc := make(map[rune]int)
	// COMMENT: there are multiple for loops here, you only need one to count chars ;)
	// Note that you don't need rune (you will use it but you can cast it directly to string assuming st is ASCII)
	for _, char := range st {
		if unicode.IsSpace(char) {
			continue
		}
		wc[char] += 1
	}
	chars := []rune{}
	for char := range wc {
		chars = append(chars, char)
	}
	for _, char := range chars {
		// COMMENT: it is miss leading that you return "" but print the result here
		fmt.Printf("%c %d\n", char, wc[char])
	}
	return ""
}

// func Occurrences(st string) map[string]int {
// 	// using strings.Field Function
// 	input := strings.Fields(st)
// 	wc := make(map[string]int)
// 	for _, word := range input {
// 		_, matched := wc[word]
// 		if matched {
// 			wc[word] += 1
// 		} else {
// 			wc[word] = 1
// 		}
// 	}
// 	return wc
// }
