package main

import (
	"fmt"

	mymath "github.com/octopus20/golang/general-practics/packages/mymath"
	mystring "github.com/octopus20/golang/general-practics/packages/mystring"
)

func main() {
	stringarray := []string{"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "A", "W", "E", "I", "O", "U", "P", "Q", "R", "S", "T", "V", "X", "Z"}
	myarray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := mymath.Total(myarray)
	average := mymath.Average(myarray)
	stringsort := mystring.SortString(stringarray)
	fmt.Println("The total of array items=", total)
	fmt.Println("The average of array items=", average)
	fmt.Println("Orderd array:", stringsort)

	// Find Occurances

	// chars count
	st := "I like learning go and I do that every day"
	st2 := "I learn go to write go scripts in go at work where I work"
	charsoccurances := mystring.Occurrences(st)
	fmt.Println(charsoccurances)

	// Words count
	wordsoccurances := mystring.Wordsoccurrences(st)
	fmt.Println(wordsoccurances)
	wordsoccurances2 := mystring.Wordsoccurrences(st2)
	fmt.Println(wordsoccurances2)

}
