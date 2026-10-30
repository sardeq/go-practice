package main

import (
	"fmt"
	"strings"
)

func handleChar() {
	var myString = []rune("mi bomba")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune = 's'
	fmt.Printf("\nRune  = %v", myRune)

	//Need to revise UTF and rune stuff later

	//-- string building --
	// var strSlice = []string{"g", "o", "o", "d", " ", "m", "o", "r", "n", "i", "n", "g"}
	// var catStr = ""

	// for i := range strSlice {
	// 	catStr += strSlice[i]
	// }
	// fmt.Printf("\n%v", catStr)

	//cannot modify strings once created, this is inefficient so we use strings package

	var strSlice = []string{"g", "o", "o", "d", " ", "m", "o", "r", "n", "i", "n", "g"}
	var stringBuilder strings.Builder
	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i]) // values are appended into array
	}
	var catStr = stringBuilder.String()
	fmt.Printf("\n%v", catStr)
}
