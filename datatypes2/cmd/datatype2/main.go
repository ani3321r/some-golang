package main

import (
	"fmt"
	"strings"
)

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

	var strSlice = []string{"raiden ", "is ", "god "}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i] //here every time a new string is being created thats why this is slower
	} //one way of concatinating strings
	fmt.Printf("\n%v", catStr)

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) //values are appended when calling this method
	} //using string builder
	var catStr1 = strBuilder.String() //here the new string is being created overall, so it is much faster
	fmt.Printf("\n%v", catStr1)
}
