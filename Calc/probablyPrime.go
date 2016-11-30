package Calc

import (
	"fmt"
	"math/big"
)

func init() {
	probablyPrime()
}
func probablyPrime() {
	i := big.NewInt(2)
	isPrime := i.ProbablyPrime(10)
	fmt.Println(isPrime)
}
