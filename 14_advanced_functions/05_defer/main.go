package main

import "fmt"

func funcOne() {
	fmt.Println("Function One")
}

func funcTwo() {
	fmt.Println("Function Two")
}

func isStudentApproved(nota1, nota2 float32) bool {
	fmt.Println("Calculando a média.")
	defer fmt.Println("Média calculada, retornando resultado.")
	media := nota1 + nota2/2
	if media <= 6 {
		return false
	}

	return true
}

func main() {
	defer funcOne() // Adiar até o último momento possível.
	funcTwo()
	fmt.Println(isStudentApproved(7, 80))
}
