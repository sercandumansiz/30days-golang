package main

import "fmt"

func main() {
	const Pi = 3.1416

	const (
		Yes = true
		No  = !Yes
	)

	const Number int8 = 3

	fmt.Println(No)
	fmt.Println(Yes)
	fmt.Println(Number)
}
