package Calc

import "fmt"

//SolveMin find min number in array
func SolveMin() {
	var n uint
	var t uint
	var i uint
	var arr []int
	fmt.Print("Enter size of array: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Print("Array value: ")
		fmt.Scan(&t)
		if t < 0 {
			err := fmt.Errorf("value %d must be positive, array type is %T\n", n, n)
			fmt.Println(err)
			break
		} else {
			arr = append(arr, int(t))
		}
	}
	minInt(arr)

}

func minInt(arr []int) {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Printf("Minimum value in array is %d\n", min)
}
