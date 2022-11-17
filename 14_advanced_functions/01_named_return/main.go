package main

import "fmt"

func main() {
	sum, sub := sumAndSub(1, 2)
	fmt.Println(sum, sub)
}

func sumAndSub(num1, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return sum, sub
}
