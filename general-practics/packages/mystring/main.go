package mystring

import (
	"sort"
	"unicode"
)

// Sort

func SortString(stringarray []string) []string {
	{
		sort.Strings(stringarray)
		return stringarray
	}
}

func Occurrences(st string) map[string]int {
	wc := make(map[rune]int)
	var result map[string]int
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
		char1 := string(char)
		result := make(map[string]int)
		result[char1] = wc[char]
		//fmt.Printf("%c %d\n", char, wc[char])
		return result
	}
	return result
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
