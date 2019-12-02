package main

import "fmt"

func main() {
	if true && true {
		fmt.Println(true)
	}

	if true && false {

	} else {
		fmt.Println(false)
	}

	if true || true {
		fmt.Println(true)
	}

	if true || false {
		fmt.Println(true)
	}

	if false || false {

	} else {
		fmt.Println(true)
	}

	if !true {
		fmt.Println(false)
	}

	if !false {
		fmt.Println(true)
	}
}
