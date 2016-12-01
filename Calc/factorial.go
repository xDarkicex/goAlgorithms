package Calc

import "fmt"

//SolveFactorial Function solves factorials
func SolveFactorial() {
	fmt.Println("Function solves for factorials ")
	var a uint
	fmt.Print("Enter Number: ")
	fmt.Scan(&a)
	f := factorial(uint(a))
	fmt.Printf("The factorial of %d is %d\n", a, f)
}

func factorial(a uint) uint {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}

