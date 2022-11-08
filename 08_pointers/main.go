// variavel que salva endereco de memoria de alguma coisa: referencia de memoria.

package main

import "fmt"

func main() {
	fmt.Println("# Pointers #")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)
	var1++
	fmt.Println(var1, var2)

	var var3 int = 100
	var pointer1 *int
	fmt.Println(var3, pointer1)

	pointer1 = &var3
	fmt.Println(var3, pointer1, *pointer1)
	var3 = 200
	fmt.Println(var3, *pointer1)
}
