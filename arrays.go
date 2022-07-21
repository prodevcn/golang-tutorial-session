package main

import "fmt"

func main() {
	fmt.Println("- Arrays -")

	var balances = [5]float32{100.00, 2.0, 3.4, 7.0, 50.0}
	var scores = []float32{100.00, 2.0, 3.4, 7.0, 50.0}
	
	balances[4] = 40.0

	fmt.Println(balances)
	fmt.Println(scores)

	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}