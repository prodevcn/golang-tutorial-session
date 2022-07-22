package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	// if
	if a < 20 {
		fmt.Println("a is less than 20\n")
	}

	// nested if
	if a == 100 {
		if b == 200 {
			fmt.Printf("Value of a is 100 and b is 200 \n")
		}
	}

	// switch
	var grade string = "B"
	marks := 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A": fmt.Printf("Excellent! \n")
	case grade == "B": fmt.Printf("Well done \n")
	case grade == "C": fmt.Printf("You passed \n")
	case grade == "D": fmt.Printf("Bettr try again \n")
	default: fmt.Printf("Invalid grade \n")
	}

	// type swtich

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x :%T \n", i)
	case int:
		fmt.Printf("type of x :%T \n", i)
	case float64:
		fmt.Printf("type of x :%T \n", i)
	case func(int) float64:
		fmt.Printf("type of x :%T \n", i)
	default:
		fmt.Printf("don't know the type")
	}

}
