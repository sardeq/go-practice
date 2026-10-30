package main

// type gasEngine struct {
// 	mpg     uint8
// 	gallons uint8
// 	owner
// }

// type owner struct {
// 	name string
// }

// func runStructs() {
// 	var myEngine gasEngine = gasEngine{25, 80, owner{"Sadeq"}}
// 	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name) // can use .name subfield directly
// }

// // anonymous structs
// type engine struct {
// 	mpg     uint8
// 	gallons uint8
// }

// func runAnoStruct() {
// 	var engi engine = engine{25, 80}

// 	//this is no reusable
// 	var engi2  = struct{
// 		mpg uint8
// 		gallons uint8
// 	}{25,90}
// 	fmt.Println(engi.mpg, engi.gallons)
// }

// type gasEngine struct {
// 	mpg     uint8
// 	gallons uint8
// }

// func (e gasEngine) milesLeft() uint8 { // method returns uint8
// 	return e.gallons * e.mpg
// }

// func runStructs() {
// 	var myEngine gasEngine = gasEngine{25, 80}
// 	fmt.Printf("Total miles left: %v\n", myEngine.milesLeft())
// }

//continue in interface.go
