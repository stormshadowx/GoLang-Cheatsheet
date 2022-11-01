package main

import "fmt"

func main() {

	// ? Declaring Variable -> Can not re-type for variable
	// ? strings
	var nameOne string = "Joe"
	var nameTwo = "stormshadow"
	var nameThree string
	nameFour := "Jerry"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)


	// ? ints
	var ageOne int = 20
	var ageTwo = 22 
	ageThree := 24
	fmt.Println(ageOne, ageTwo, ageThree)

	// ? bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 10 // uint for real number only
	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 999343993202.2233
	scoreThree := 20302.21
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
