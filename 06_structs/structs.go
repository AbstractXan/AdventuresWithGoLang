package main

import "fmt"

//Struct
// Variable with first letter capital can be used outside package
// Variable with first letter small could only be used inside the same package

type Example struct {
	Val   string
	count int
}

func main() {
	var ex Example

	//Unlike C++ , new struct instances can be intitialized partially as follows:
	ex.Val = "Hello"
	
	fmt.Println(ex)
}

// Output: {Hello 0}
