package main

import "fmt"

/***
	Methods are just like functions except that there is one extra parameter - the receiver.
	Receiver is declared before the function
***/
//Custom type
type kilometer int

//Method for cutom type
func (l kilometer) PrintMiles() {
	miles := float64(l) * 0.62
	fmt.Println(l, "km is equal to ", miles, "miles.")
}

//Custom struct
type Example struct {
	Val   string
	count int
}

//Method for struct
//Note capital L implies that this method could be used outside of this package
//The method on the structure in the example at the start of this section takes a pointer as the receiver.
//This means that it can modify fields of the receiver and these changes will be shared.

func (e *Example) Log() {
	e.count++
	fmt.Println(e.count, e.Val)
}

func main() {

	//Custom type
	var length kilometer = 1
	length.PrintMiles()

	/***
		Note that one can assign integer to kilometer type but not the other way.
		This proves really effective in cases where we wanna define variables like kilmetres and miles.
		And both shouldn't be assigned to one another.
	***/

	//Custom struct
	example := Example{"Hi", 2}
	example.Log()

}
