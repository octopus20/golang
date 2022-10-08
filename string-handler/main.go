package main

import "fmt"

var equation int

// basic multiplex function , that asks the user to enter two numbers and show the result
func main() {
	var first, second int
	fmt.Print("Enter first number : ")
	fmt.Scan(&first) // fmt.Scan will read the user input, from the pointer
	fmt.Print("Enter second number : ")
	fmt.Scan(&second)
	equation = first * second
	fmt.Println("The result is: ", equation)
}
