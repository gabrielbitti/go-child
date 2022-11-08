package main

import (
	"errors"
	"fmt"
)

func main() {
	// Int
	fmt.Println("# Int:")
	var num1 int = -1000000000
	var num2 uint32 = 2000
	fmt.Println(num1)
	fmt.Println(num2)

	// Float
	fmt.Println("# Float:")
	var real1 float32 = 123.45
	var real2 float64 = 121231231231.23
	fmt.Println(real1)
	fmt.Println(real2)

	// Strings
	fmt.Println("# Strings:")
	var str1 string = "String #1"
	str2 := "String #2"
	fmt.Println(str1)
	fmt.Println(str2)

	// Char
	// Represents the value in a number of ASCII table
	fmt.Println("# Char: represents the value in a number of ASCII table")
	char1 := '%'
	fmt.Println("% -> ", char1)

	// Boolean
	fmt.Println("# Boolean:")
	var bool1 bool = true
	var bool2 bool
	fmt.Println(bool1)
	fmt.Println(bool2)

	// Error
	fmt.Println("# Error Type:")
	var err1 error
	err2 := errors.New("internal error: it's not a string")
	fmt.Println(err1)
	fmt.Println(err2)

	// Aliases
	fmt.Println("# Aliases:")

	// Rune: equal to int32
	var num3 rune = 123456
	fmt.Println("## Rune: ", num3)

	// Byte: equal to int8
	var num4 byte = 123
	fmt.Println("## Byte: ", num4)
}
