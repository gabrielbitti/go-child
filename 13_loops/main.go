package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("# Loops")

	i := 0
	fmt.Println("Value of i: ", i)

	for i < 3 {
		i++
		fmt.Println(">>> Incrementing var i: ", i)
		time.Sleep(time.Second)
	}

	fmt.Println("Result of i: ", i)
	fmt.Println("====================")

	for j := 0; j <= 6; j += 2 {
		fmt.Println(">>> Incrementing var j: ", j)
		time.Sleep(time.Second)
	}

	users := [3]string{"John", "Marie", "Heisenberg"}
	for index, value := range users {
		fmt.Println("Index, value: ", index, value)
	}

	fmt.Println("====================")

	for _, value := range users {
		fmt.Println(value)
	}

	fmt.Println("====================")

	for _, word := range "WORD" {
		fmt.Println(string(word))
	}

	users2 := map[string]string{
		"first_name": "Leon",
		"last_name":  "Rooster",
	}

	fmt.Println("====================")

	for index, value := range users2 {
		fmt.Println(index, value)
	}
}
