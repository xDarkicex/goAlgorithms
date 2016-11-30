package Calc

import (
	"fmt"
	"math/big"
)

//GuessPrime guess is very large numbers are prime
func GuessPrime() {
	fmt.Println("Guess if very large number is prime. ")
	var a int64
	fmt.Print("Enter Number: ")
	_, _ = fmt.Scan(&a)
	probablyPrime(a)
}
func probablyPrime(a int64) {
	i := big.NewInt(a)
	isPrime := i.ProbablyPrime(200)
	if isPrime {
		fmt.Printf("%d probably is prime.\n", i)
	} else {
		fmt.Printf("%d probably isn't prime.\n", i)
	}
}
