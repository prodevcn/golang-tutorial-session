package main

import "fmt"

func main() {
	fmt.Println("- Recursion -");

	i:= 15
	fmt.Printf("Factorial of %d is %d \n", i, factorial(i))

	for j := 0; j < 10; j++ {
		fmt.Printf("%d", fibonaci(j))
	}
}

func factorial(i int) int {
	if (i <= 1) {
		return 1
	}

	return i * factorial(i - 1)
}

func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return fibonaci(i - 1) + fibonaci(i - 2)
}
