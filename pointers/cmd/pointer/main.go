package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("p points to: %v", *p)
	fmt.Printf("\nvalue of i: %v", i)
	p = &i //p and i both references the same int32 value
	*p = 50
	fmt.Printf("\np points to: %v", *p)
	fmt.Printf("\nvalue of i: %v", i)
	var k int32 = 23
	i = k //copy the value of k to i's memory location

	var num = [5]float64{2, 4, 6, 8, 10}
	fmt.Printf("\nmemory location of num array: %p", &num)
	var res [5]float64 = square(num)
	fmt.Printf("\nthe result is: %v", res)
	fmt.Printf("\nthe result is: %v", num)
}

func square(num1 [5]float64) [5]float64 {
	for i := range num1 {
		num1[i] = num1[i] * num1[i]
	}
	return num1
}

//by using the pointer we are saving space as well as pointing both the variables to one array
//we don't need to copy data, so it saves time as well as memory

// func square(num1 *[5]float64) [5]float64 {
// 	for i := range num1 {
// 		num1[i] = num1[i] * num1[i]
// 	}
// 	return num1
// }

// var res [5]float64 = square(&num)
