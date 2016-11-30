package Calc

import "fmt"

//SolveMax find min number in array
func SolveMax() {
	var n int64
	var t int64
	var i int64
	var arr []int
	fmt.Print("Enter size of array: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Print("Array value: ")
		fmt.Scan(&t)
		if t < 0 {
			err := fmt.Errorf("value %d must be positive, array type is %T\n", t, n)
			fmt.Println(err)
			break
		} else {
			arr = append(arr, int(t))
		}
	}
	maxInt(arr)

}

func maxInt(arr []int) {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Printf("Maximum value in array is %d\n", max)
}
