package Calc

import "fmt"

//Total Calc array total
func Total() {
	var n int
	var t uint
	fmt.Print("Enter size of Array")
	fmt.Scan(&n)
	var sum uint
	for i := 0; i < n; i++ {
		fmt.Print("Array value: ")
		fmt.Scan(&t)
		sum += t
	}
	fmt.Printf("Array Total [%d]", sum)
}
