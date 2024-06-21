package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_FISH_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var fishChannel = make(chan string)
	var shops = []string{"spencer", "licious", "local"}
	for i := range shops {
		go checkChickenPrices(shops[i], chickenChannel)
		go checkFishPrices(shops[i], fishChannel)
	}
	sendMessage(chickenChannel, fishChannel)
}

func checkFishPrices(shop string, fishChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var fishPrice = rand.Float32() * 20
		if fishPrice <= MAX_FISH_PRICE {
			fishChannel <- shop
			break
		}
	}
}

func checkChickenPrices(shop string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- shop
			break
		}
	}
}

func sendMessage(chickenChannel chan string, fishChannnel chan string) {
	select {
	case shop := <-chickenChannel:
		fmt.Printf("\nDeal on chicken at %v", shop)
	case shop := <-fishChannnel:
		fmt.Printf("\nDeal on fish at %v", shop)
	} //works like a switch case
}
