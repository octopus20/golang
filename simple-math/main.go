package main

import "fmt"

var equation int

func main() {
	var first, second int
	fmt.Print("Enter first number : ")
	fmt.Scan(&first) // fmt.Scan will read the user input
	fmt.Print("Enter second number : ")
	fmt.Scan(&second)
	equation = first * second
	fmt.Println("The result is: ", equation)
}
