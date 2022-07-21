package main

import "fmt"

func main() {
	fmt.Println("- Functions -")
	x := 10
	y := 20

	z := max(x, y)
	fmt.Printf("Max value is : %d\n", z);

	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

func max(num1 int, num2 int) int {
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}

	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
