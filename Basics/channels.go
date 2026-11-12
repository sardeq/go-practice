package main

import (
	"fmt"
	"time"
)

func chn() {
	//var c = make(chan int) //normal channel
	var c = make(chan int, 5) //buffer channel

	// c <- 1
	// var i = <-c
	// fmt.Println(i)
	// this will cause deadlock aka waiting forever so we have to use channels with goroutines

	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

}

func process(c chan int) {
	defer close(c) //mean close right before the func exits
	for i := range 5 {
		c <- i
	}
	fmt.Println("Exiting process")
}
