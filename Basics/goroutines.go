package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex prevent race conditions when multiple goroutines access shared dat
// var m = sync.Mutex{}      //right now we have unexpected results so we need a mutex to lock and unlock each part
var m = sync.RWMutex{} //this is the same but allows multiple goroutines to read shared data simultaneously, but only one goroutine to write at a time, blocking all readers while writing.

var wg = sync.WaitGroup{} //pretty much a counter
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func rnRt() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    //whenever we spawn a go routine we add 1
		go dbCall(i) //makes it run concurrently
	}
	wg.Wait() //wait for counter to be 0 again
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done() //decrements the counter
}

func save(result string) {
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The results are: %v", results)
	m.RUnlock()
}
