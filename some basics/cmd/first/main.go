package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!!")
	var intNum int = 2000
	fmt.Println(intNum)

	var myRune rune = 'a'
	fmt.Println(myRune)

	fmt.Println(len("γ"))                    //here the output will be 2
	fmt.Println(utf8.RuneCountInString("γ")) //using this we get the exact count of special chars

	var myVar = "hello" //automatically recognizes the type
	NewVar := "hi"      //this is also a way to initialize a variable
	fmt.Println(myVar)
	fmt.Println(NewVar)

	var1, var2 := 5, 6 //we can also initialize two variables at  once
	fmt.Println(var1, var2)

	const pi float32 = 3.1415
	fmt.Println(pi)
}
