package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

func(circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	fmt.Println("- Functions -")
	x := 10
	y := 20

	z := max(x, y)
	fmt.Printf("Max value is : %d\n", z);

	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)

	// call function by value
	fmt.Printf("Before swap_value, value of x: %d \n", x)
	fmt.Printf("Before swap_value, value of y: %d \n", y)

	swap_value(x, y)

	fmt.Printf("Before swap_value, value of x: %d \n", x)
	fmt.Printf("Before swap_value, value of y: %d \n", y)


	// call function by reference
	fmt.Printf("Before swap_by_ref, value of x : %d\n", x)
	fmt.Printf("Before swap_by_ref, value of y : %d\n", y)

	swap_by_ref(&x, &y)

	fmt.Printf("After swap_by_ref, value of x : %d\n", x)
	fmt.Printf("After swap_by_ref, value of y : %d\n", y)

	// function as value
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	// function as closure
	nextNumber := getSequence()

	/* invoke nextNumber to increase i by 1 and return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	
	/* create a new sequence and see the result, i is 0 again*/
	nextNumber1 := getSequence()  
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	// call function as method
	circle := Circle {x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())
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

// call by value
func swap_value(x, y int) int {
	var temp int

	temp = x
	x = y
	y = temp
	
	return temp
}


// call by reference
func swap_by_ref(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

// function closure
func getSequence() func () int {
	i := 0
	return func () int {
		i ++
		return i
	}
}