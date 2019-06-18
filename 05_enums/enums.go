package main

import "fmt"

func main() {
	const (
		Red   = iota
		Blue  = iota
		Green = iota
	)

	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)
}
