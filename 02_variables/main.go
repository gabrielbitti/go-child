package main

import "fmt"

func main() {
	var var1 string = "Var #1"
	var2 := "Var #2"
	fmt.Println(var1, var2)

	var (
		var3 string = "Var #3"
		var4 string = "Var #4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "Var #5", "Var #6"
	fmt.Println(var5, var6)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)

	const const1 = "Const #1"
	fmt.Println(const1)
}
