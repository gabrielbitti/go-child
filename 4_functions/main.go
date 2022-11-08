package main

import "fmt"

func main() {
	result := sum(2, 1)
	fmt.Println(result)

	varF := func(txt string) {
		fmt.Println("Function F: ", txt)
	}

	varF("Testing")

	sumResult, subResult := mathCalc(10, 15)
	fmt.Println(sumResult, subResult)

	sumResult2, _ := mathCalc(10, 15)
	fmt.Println(sumResult2)
}

func sum(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func mathCalc(num1, num2 int8) (int8, int8) {
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}
