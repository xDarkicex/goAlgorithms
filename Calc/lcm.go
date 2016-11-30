package Calc

import "fmt"

// LCM Least common Multiple
func LCM(a, b uint) uint {
	return (a / GCD(a, b)) * b
}

// SolveLCM Find least common Multiple
func SolveLCM() {
	fmt.Println("LCM takes two numbers ")
	var a, b uint
	fmt.Print("Enter Number: ")
	_, _ = fmt.Scan(&a)
	fmt.Print("Enter Number: ")
	_, _ = fmt.Scan(&b)
	LCM := LCM(uint(a), uint(b))
	fmt.Printf("The Least common Multiple for %d and %d is %d\n", a, b, LCM)
}
