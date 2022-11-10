package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada com sucesso!!!")
	}
}

func isStudentApproved(nota1, nota2 float32) bool {
	defer recoverExecution()
	media := (nota1 + nota2) / 2
	if media < 6 {
		return false
	} else if media > 6 {
		return true
	}

	panic("Média é exatamente 6.")
}

func main() {
	//fmt.Println(isStudentApproved(6, 7))
	fmt.Println(isStudentApproved(6, 6))
	fmt.Println("Após a execução.")
}
