package main

import (
	"fmt"
)

func handleMaps() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{
		"Adam":  27,
		"Sadeq": 67,
	}
	fmt.Println(myMap2["Sadeq"])

	delete(myMap2, "Adam")

	var age, ok = myMap2["Tung"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid")
	}

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// }

	// for i := range 10 {
	// 	fmt.Println(i)
	// }

}
