package main

import "fmt"

func main() {
	fmt.Println("# Variations Functions")
	result := sumNumbers(10, 20, 300)
	fmt.Println(result)

	writeText("Hello, World", 1, 2, 3, 50)
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func writeText(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}
