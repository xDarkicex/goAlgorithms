//sieve of eratosthenes algorithm

package Calc

import "fmt"

// SolveSieve passes prime function number
func SolveSieve() {
	fmt.Println("Find prime numbers through nth number")
	var a uint
	fmt.Print("Enter Number: ")
	_, _ = fmt.Scan(&a)
	primes := Sieve(a)
	for _, prime := range primes {
		fmt.Printf("%d\n", prime)
	}
	fmt.Printf("\n")
}

// Sieve solves primes
func Sieve(N uint) (primes []uint) {
	b := make([]bool, N)
	var i uint
	for i = 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}
