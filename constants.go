package main

import "fmt"

func main() {
	fmt.Println("- Constant value -")

	// integer literals
	x := 212
	fmt.Printf("x is type of %T\n", x)

	y := 0x4b
	fmt.Printf("y is type of %T\n", y)
	fmt.Println(y);

	// const keyword
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d", area);

}