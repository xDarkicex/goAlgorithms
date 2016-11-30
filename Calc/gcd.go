package Calc

import "fmt"

//GCD greatest common demonmiator
func GCD(a uint, b uint) uint {
	if b == 0 {
		return a
	}
	r := a % b
	return GCD(b, r)
}

// SolveGCD sets up inputs might at parse flags if I use this a lot
func SolveGCD() {
	fmt.Println("GCD takes two numbers ")
	var a, b uint
	fmt.Print("Enter Number: ")
	_, _ = fmt.Scan(&a)
	fmt.Print("Enter Number: ")
	_, _ = fmt.Scan(&b)
	GCD := GCD(uint(a), uint(b))
	fmt.Printf("The Greatest common Denominator for %d and %d is %d\n", a, b, GCD)
}
