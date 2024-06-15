package main

import "fmt"

func main() {
	var newString = "l'accent tréma"
	fmt.Println(newString)
	var indexed = newString[0]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range newString {
		fmt.Println(i, v)
	} //many special characters are missed because not support it by default

	var string2 = []rune("l'accent tréma") //here comes rune, which supports all characters as it has the capacity to store binary worth int32
	var indexed1 = string2[0]
	fmt.Printf("%v, %T", indexed1, indexed1)

}
