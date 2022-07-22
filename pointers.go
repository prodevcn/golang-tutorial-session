package main

import "fmt"

const MAX int = 3

func main() {
	fmt.Println("- Pointers -")

	var a int = 20
	var ip *int
	
	ip = &a

	fmt.Printf("Address of a variable: %x\n", &a)
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)

	// array of pointer
	score := []int{10, 100, 200}
	var i int
	var ptrs [MAX]*int

	for i = 0; i < MAX; i++ {
		ptrs[i] = &score[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptrs[i])
	}

	// a pointer to a pointer
	var x int
	var pointer *int
	var pointter **int

	x = 3000

	pointer = &x
	pointter = &pointer

	fmt.Printf("Value of x = %d\n", x )
	fmt.Printf("Value available at *pointer = %d\n", *pointer )
	fmt.Printf("Value available at **pointter = %d\n", **pointter)

	// passing pointers to function
	var m int = 100
	var n int = 200

	swap(&m, &n)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}