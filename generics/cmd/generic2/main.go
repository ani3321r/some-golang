package main

import "fmt"

type gasEngine struct {
	liter float32
	kmpl  float32
}

type electricEngine struct {
	kwh    float32
	kmpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "BMW",
		carModel: "M5",
		engine: gasEngine{
			liter: 60,
			kmpl:  7.25,
		},
	}
	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "cybertruck",
		engine: electricEngine{
			kwh:    122.85,
			kmpkwh: 3.22,
		},
	}
	fmt.Printf(gasCar.carMake)
	fmt.Printf(electricCar.carMake)
}
