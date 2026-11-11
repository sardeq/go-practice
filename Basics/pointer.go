package main

import "fmt"

func point() {
	var p *int32 = new(int32) //it initialized 8 bytes in memory without new its nill, with new its the address for this pointer
	var i int32
	fmt.Printf("the value p points to is %v\n", *p)
	fmt.Printf("i=%v\n", i)
	// *p = 10 give the value p references a value
	p = &i //refernces memory address of i
	*p = 8
	fmt.Printf("the value p points to is %v\n", *p)
	fmt.Printf("i=%v\n", i)
	//this is not like int k = i, because it copies the values to a new memory address
}

func slicePoint() {
	//uh so copying slices isnt like variables, they point to same memory address (if i understood correctly)
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}

func pointExample() {
	var y = [5]float64{2, 6, 3, 8}
	fmt.Printf("memeory location of y array is %p\n", &y)
	var result [5]float64 = square(&y)
	fmt.Printf("The result is %v\n", result)

	//we pass pointers so we dont use double the memory, but if we need to access different variables i think its better to just do it normally
	// my memory isnt that great but when i was learning C we can still handle a case like this without using double the memory ill look into it later
}

func square(x *[5]float64) [5]float64 {
	fmt.Printf("memeory location of x array is %p\n", &x)
	for i := range x {
		x[i] = x[i] * x[i]
	}
	return *x
}
