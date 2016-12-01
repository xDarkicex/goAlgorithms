package Calc

import "fmt"

var count uint64

//SolveCollatz ...
func SolveCollatz() {
	fmt.Println("Function solves for Collatz ")
	var a uint
	fmt.Print("Enter Number: ")
	fmt.Scan(&a)
	f := collatz(uint64(a))
	fmt.Printf("The collatz number of %d is %d\n", a, f)
}

func collatz(a uint64) uint64 {
	fmt.Printf("Step: %d\n", a)
	if a <= 0 {
		fmt.Printf("Hit %d. Not supposed to happen\n", a)
		return count
	} else if a == 1 {
		return count
	}
	count++
	if a%2 == 0 {
		return collatz(a / 2)
	}
	return collatz(a*3 + 1)
}
