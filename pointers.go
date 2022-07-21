package main

import "fmt"

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

}