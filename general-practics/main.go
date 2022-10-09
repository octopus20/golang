package main

import (
	"fmt"

	mymath "github.com/octopus20/golang/general-practics/mymath"
	mystring "github.com/octopus20/golang/general-practics/mystring"
)

func main() {
	stringarray := []string{"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "A", "W", "E", "I", "O", "U", "P", "Q", "R", "S", "T", "V", "X", "Z"}
	myarray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := mymath.Total(myarray)
	average := mymath.Average(myarray)
	stringsort := mystring.SortString(stringarray)
	fmt.Println(total)
	fmt.Println(average)
	fmt.Println(stringsort)

}
