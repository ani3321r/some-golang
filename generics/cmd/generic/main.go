package main

import "fmt"

func main() {
	var intSlice = []int{4, 9, 7}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{3, 2, 1}
	fmt.Println(sumSlice[float32](float32Slice))
	fmt.Println(isEmpty(float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
} //decalring many types to one function

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
} //type "any" cannot be assigned to all types of functions, only used when returning value independent things
