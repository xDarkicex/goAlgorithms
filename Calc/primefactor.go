package Calc

import "fmt"

//SolveFactor ...
func SolveFactor() {
	fmt.Println("Function solves for factors.")
	var a uint
	fmt.Print("Enter Number: ")
	fmt.Scan(&a)
	f := factor(uint(a))
	fmt.Printf("The factors of %d is %d\n", a, f)
}

var factored []uint

func factor(a uint) []uint {
	primes := Sieve(a)
	for _, k := range primes {
		if a%k == 0 {
			factored = append(factored, k)
			return factor(a / k)
		}
	}
	return append(factored, a)
}
