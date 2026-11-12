package main

import "fmt"

func runGeneric() {
	var intSlice = []int{}
	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))
}

func sumSlice[T int | float32 | float64](slice []T) T { //can use any instead of typing seperately but doesnt work in this example
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
