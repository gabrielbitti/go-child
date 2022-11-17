package main

import "fmt"

func closure() func() {
	text := "Inside closure function."
	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	// referenciam variaveis fora do corpo
	fmt.Println("# 14.07 Closure functions")

	text := "Inside main function."
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}
