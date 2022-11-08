package main

import "fmt"

func main() {
	fmt.Println("# Control structure")
	number := 10
	if number > 15 {
		fmt.Println("Major than 15")
	} else {
		fmt.Println("Minor than 15")
	}

	if outro := number; outro > 0 {
		fmt.Println("Major than 0")
	}
}
