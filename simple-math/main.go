package main

import (
	"fmt"
	"log"
)

var equation int

func main() {
	var first, second int
	fmt.Print("Enter first number : ")
	// Scan and validate user input
	if _, err := fmt.Scan(&first); err != nil {
		log.Print("  Scan for first input failed, due to ", err)
		return
	}
	fmt.Print("Enter second number : ")
	// Scan and validate user input
	if _, err := fmt.Scan(&second); err != nil {
		log.Print("  Scan for second input failed, due to ", err)
		return
	}
	equation = first * second
	fmt.Println("The result is: ", equation)
}
