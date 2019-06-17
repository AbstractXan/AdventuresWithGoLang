package main

import "fmt"

//Function parameters (usually of same type)
//Variadic function (returns multiple parameters)
func add(a, b int) (string, int) {
	c := a + b
	return "No error", c
}
func main() {

	//Blank identifier '_' to discard unused values
	_, sum := add(3, 4)
	fmt.Println(sum)

	// Closure functions
	// Written inside functions
	closure := func(msg string) {
		sum++
		fmt.Println(sum, msg)
	}

	//Calling closure functions
	closure("First")
	closure("Second")
}
