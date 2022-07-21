package main

import "fmt"

func main() {
	fmt.Println("- Slice -")

	var numbers = make([]int, 3, 5)
	var scores []int
	// fmt.Println(number)

	printSlice(numbers)

	if(scores == nil) {
		fmt.Printf("slice is nil\n")
	}

	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(values)

	/* print the original slice */
	fmt.Println("numbers ==", values)
   
	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] ==", values[1:4])
	
	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", values[:3])
	
	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", values[4:])
	
	numbers1 := make([]int,0,5)
	printSlice(numbers1)
	
	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := values[:2]
	fmt.Println(number2)
	printSlice(number2)
	
	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := values[2:5]
	fmt.Println(number3)
	printSlice(number3)
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}