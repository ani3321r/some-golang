package main

import "fmt"

type laptop struct {
	power   uint8
	ram     uint64
	display //using another struct
}

type desktop struct {
	power uint16
	ram   uint64
}

type display struct {
	displayType string
}

func (e laptop) ram_in_laptop() uint64 {
	return e.ram * 1024
} //inheritence

func (e desktop) ram_in_desktop() uint64 {
	return e.ram * 1024
} //inheritence

type ram_in_mb interface {
	ram_in_laptop() uint64
} //instead of using two different structs we can use just a interface

func can_run(e ram_in_mb, totalMem uint64) {
	if totalMem >= e.ram_in_laptop() {
		fmt.Println("You need more ram")
	} else {
		fmt.Println("You can run the program")
	}
}

func main() {
	var mac laptop = laptop{60, 16, display{"xdr"}}
	fmt.Println(mac.power, mac.ram, mac.displayType)

	var legion5i laptop = laptop{90, 16, display{"oled"}}
	can_run(legion5i, 16000)

	//var alienware desktop = desktop{450, 32}
	//can_run(alienware, 30000)
}
