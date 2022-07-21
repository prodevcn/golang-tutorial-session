package main

import "fmt"

func main() {
	fmt.Println("- Operators -")

	// bitwise operators
	a := 60
	b := 30
	c := a & b
	d := a ^ b
	e := a | b
	f := a >> 2
	g := a << 2
	

	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", e)
	fmt.Printf("%d\n", f)
	fmt.Printf("%d\n", g)

	// address and reference operators
	x := 20
	y := &x
	z := *y
	fmt.Printf("%d", y)
	fmt.Printf("%d", z)
}