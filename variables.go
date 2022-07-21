package main

import "fmt"

func main() {
	fmt.Println("- Variables -")

	// static type declaration in Go
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	// dynamic type declaration / type inference in Go
	var y float64 = 20.0

	z := 42

	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)

	// mixed variable declaration in Go
	var a, b, c = 3, 4, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is type of %T\n", a)
	fmt.Printf("b is type of %T\n", b)
	fmt.Printf("c is type of %T\n", c)

	// lvalue and rvalue in Go
	/*
		lvalue - Expressions that refer to a memory location is called "lvalue" expression. An lvalue may appear as either the left-had or right-hand side of an assignment.
		rvalue - The term rvalue refers to a data value that is stored at some address in memory. An rvalue is an expression that cannot have a value assigned to it which means an rvalue may appear on the right -but not left-had side of an assignment.
	*/

	x = 20.0
}
