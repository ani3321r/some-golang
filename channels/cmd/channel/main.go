package main

import "fmt"

func main() {
	var c = make(chan int)
	go process(c)
	for i := 0; i < 6; i++ {
		fmt.Println(<-c) //directly accessing the value in the channel
	}
	// for i := range c {
	// 	fmt.Println(i)
	// }this will give a deadlock error as c will return to the loop and seaech for the next call

	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c) //due to this we can use the range, as it works like break
	for i := 0; i < 6; i++ {
		c <- i //assigning a value inside a channel
	}
}
