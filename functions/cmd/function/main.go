package main

import (
	"errors"
	"fmt"
)

func main() {
	var value string = "raiden"
	printme(value)
	var val1 int = 11
	var val2 int = 0
	var result, remainder, err = division(val1, val2)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the divison is %v", result)
	} else {
		fmt.Printf("the remainder is %v and result after division is %v", remainder, result)
	}
}

func printme(value string) {
	fmt.Println(value)
}

func division(val1 int, val2 int) (int, int, error) {
	var err error
	if val2 == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = val1 / val2
	var remainder int = val1 % val2
	return result, remainder, err
}
