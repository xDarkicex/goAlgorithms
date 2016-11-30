package Calc

import "fmt"

// Variance find diff in array
func Variance() {
	var arr []float64
	var n float64
	var t float64
	var i float64
	fmt.Print("Enter size of array: ")
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Print("Array value: ")
		fmt.Scan(&t)
		if t < 0 {
			err := fmt.Errorf("value %f must be positive, array type is %T\n", t, n)
			fmt.Println(err)
			break
		} else {
			arr = append(arr, t)
		}
	}
	s := total(arr)
	fmt.Println("Total: ", s)
	mean := s / n
	fmt.Println("Average: ", mean)
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += (arr[i] - mean) * (arr[i] - mean)
	}
	v := sum / n
	fmt.Printf("variance of array %f\n", v)
}
func total(arr []float64) float64 {
	var sum float64

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum

}
