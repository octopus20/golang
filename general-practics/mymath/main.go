package mymath

func Total(myarray []int) int {
	sum := 0
	for i := 0; i < (len(myarray)); i++ {
		sum += (myarray[i])
	}
	return sum
}
func Average(myarray []int) float64 {
	sum := 0
	for i := 0; i < (len(myarray)); i++ {
		sum += (myarray[i])
	}
	avg := (float64(sum)) / (float64(len(myarray)))
	return avg
}
