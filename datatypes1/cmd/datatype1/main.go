package main

import "fmt"

func main() {
	var intArr [4]int32
	intArr[1] = 32
	fmt.Println(intArr[1:4])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2]) //contiguous memory allocated

	arr1 := [3]int32{2, 4, 6} //we can also initialize array like this
	fmt.Println(arr1)

	var intSlice []int32 = []int32{8, 11, 14}
	fmt.Println(intSlice)
	fmt.Printf("the length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 16) //a new array is being created with a higher capacity to append the new value
	fmt.Printf("\n")
	fmt.Println(intSlice)
	fmt.Printf("the length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{18, 20}
	intSlice = append(intSlice, intSlice2...) //another way to
	fmt.Println(intSlice)

	//var intSlice3 []int32 = make(int32[], 4, 8)//this is how we allocate the size and capacity of a slice

	var map1 map[string]uint8 = make(map[string]uint8) //initialize a map
	fmt.Println(map1)

	var map2 = map[string]uint8{"raiden": 96, "ada": 51, "aab": 56} //assign value in map
	fmt.Println(map2["raiden"])
	fmt.Println(map2["abc"]) //it will return a default value even if key is absent
	delete(map2, "ada")      //delete a key of map

	for name, age := range map2 {
		fmt.Printf("name: %v, age: %v \n", name, age)
	}

	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	} //another way of writing for loop
}
