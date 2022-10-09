package mymath

import (
	"fmt"
)

func Main() {

	myarray := []int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < (len(myarray)); i++ {
		sum += (myarray[i])
	}
	avg := (float64(sum)) / (float64(len(myarray)))

	fmt.Println("Total = ", sum, "\nAverage = ", avg)
}
