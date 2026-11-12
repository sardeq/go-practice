package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_SHAWARMA_PRICE float32 = 5
var MAX_BORGER_PRICE float32 = 5

func shCheck() {
	var shawarmaChannel = make(chan string)
	var borgerCHannel = make(chan string)
	var websites = []string{"taroob.com", "gunes.com", "shawarmeronsaj.com"}
	for i := range websites {
		go checkShawarmaPrice(websites[i], shawarmaChannel)
		go checkBorgerPrice(websites[i], borgerCHannel)
	}
	sendMessage(shawarmaChannel, borgerCHannel)
}

func checkShawarmaPrice(website string, shawarmaChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var shawarmaPrice = rand.Float32() * 20
		if shawarmaPrice <= MAX_SHAWARMA_PRICE {
			shawarmaChannel <- website
			break
		}
	}
}

func checkBorgerPrice(website string, borgerCHannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var borgerPrice = rand.Float32() * 20
		if borgerPrice <= MAX_BORGER_PRICE {
			borgerCHannel <- website
			break
		}
	}
}

func sendMessage(shawarmaChannel chan string, borgerChannel chan string) {
	select {
	case website := <-shawarmaChannel:
		fmt.Printf("SMS: Found a deal on shawarma at %v\n", website)
	case website := <-borgerChannel:
		fmt.Printf("Email: Found a deal on borger at %v\n", website)
	}
}
