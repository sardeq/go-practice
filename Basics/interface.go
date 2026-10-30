package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkhw uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) mileLeft() uint8 {
	return e.mpkhw * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You need more fuel")
	}
}

func runStructs() {
	var myEngine gasEngine = gasEngine{25, 80} // can do electirc engine as well and it will work due to interface
	canMakeIt(myEngine, 50)
}
