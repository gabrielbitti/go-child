package main

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func invertPointerSignal(number *int) {
	*number = *number * -1
}

func main() {
	// referenciam variaveis fora do corpo
	fmt.Println("# 14.08 Functions with pointers")

	number := 20
	invertedNumber := invertSignal(number)
	fmt.Println(invertedNumber)
	fmt.Println(number)

	fmt.Println("===========")

	newNumber := 30
	fmt.Println(newNumber)
	invertPointerSignal(&newNumber)
	fmt.Println(newNumber)

}
