package main

import "fmt"

func main() {
	// Arictimetics
	sum := 1 + 1
	sub := 1 - 1
	div := 100 / 2
	multiplication := 10 * 1
	restOfDivision := 10 % 3
	fmt.Println(sum, sub, div, multiplication, restOfDivision)
	//end

	// Atribuition
	var string = "string 1"
	string2 := "string 2"
	fmt.Println(string, string2)
	// end

	// relacionais
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)

	// logicos
	truee, falsee := true, false
	fmt.Println(truee && falsee) // and
	fmt.Println(truee || falsee) // or
	fmt.Println(!truee)          // or
	fmt.Println(!falsee)         // or

	//unarios
	owrNumber := 10
	owrNumber++
	owrNumber += 15
	fmt.Println(owrNumber)

	// Ternário: nao existe
}
