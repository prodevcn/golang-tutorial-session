package main

import "fmt"

// global variable declaration
var g int = 50
var x int = 20

func main() {
	fmt.Println("- Scope Rules -")

	// local variables declaration
	var a, b, c int
	var x int = 10
	var y int = 20
	var z int = 0


	// local variable of same name with global variable
	var g int
	
	// actual intialiization
	a = 10
	b = 20
	c = a + b
	g = a - b


	fmt.Printf("value of a = %d, b = %d, c = %d and g = %d\n", a, b, c, g)

	fmt.Printf("value of x in main() = %d\n", x)
	z = sum(x, y)
	fmt.Printf("value of z in main() = %d\n", z)

}

func sum(x, y int) int {
	fmt.Printf("value of x in sum() = %d\n", x)
	fmt.Printf("value of y in sum() = %d\n", y)

	return x + y
}