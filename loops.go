package main

import "fmt"

func main() {
	fmt.Println("- Loops -")

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}
	
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d \n", a)
	}

	for a < b {
		a++
		fmt.Printf("value of a: %d \n", a)
	}

	for i,x := range numbers {
		fmt.Printf("value of x = %d at %d \n", x, i)
	}

	// nested loop
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i/j); j++ {
			if (i%j == 0) {
				break;
			}
		}

		if(j > (i/j)) {
			fmt.Printf("%d is prime \n", i)
		}
	}

	// goto
	var x int = 10
	LOOP: for x < 20 {
		if x == 15 {
			x += 1
			goto LOOP
		}

		fmt.Printf("value of x : %d \n", x)
		x++
	}
}