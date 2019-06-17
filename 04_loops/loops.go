package main

import "fmt"

//While loop until a < 0
func whileLessThanZero(a int) {
	for a < 0 {
		fmt.Println("a = ", a)
		a++
	}
}

// A typical for loop
func printFive() {
	for j := 1; j <= 5; j++ {
		fmt.Println("j = ", j)
	}
}

//Break from a labelled loop
// Here, the outer loop is labelled as 'K'
func looped() {
	for i := 5; i < 10; i++ {
	K:
		for {
			for {
				fmt.Println("i = ", i)
				break K
			}
		}

	}
}
func main() {
	whileLessThanZero(-2)
	printFive()
	looped()
}
