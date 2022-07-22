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

	// multi-dimensional array
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var k, h int

	for k = 0; k< 5; k++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d \n", k, h, a[k][h])
		}
	}

	// passing array as parameter
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	avg = getAverage(balance, len(balance))
	
	fmt.Printf("Average value is %f \n", avg)

}

// passing array as function parameter
func getAverage(arr []int, size int) float32 {
	var i int
	var avg, sum float32

	for i = 0; i< size; i++ {
		sum += float32(arr[i])
	}

	fmt.Printf("Sum inside of getAverage is : %f \n", sum)

	avg = sum / float32(size)
	return avg
}