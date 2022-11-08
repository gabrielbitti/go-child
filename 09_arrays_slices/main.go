// variavel que salva endereco de memoria de alguma coisa: referencia de memoria.

package main

import "fmt"

func main() {
	fmt.Println("# Arrays and Slices #")
	// array é uma lista de valores DO MESMO TIPO
	var array1 [5]int
	array1[2] = 3
	fmt.Println(array1)

	array2 := [5]string{"#1", "#2", "#3", "#4", "#5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	// slices -- it's not an array
	slice1 := []int{10, 2, 4, 69, 29}
	fmt.Println(slice1)

	slice1 = append(slice1, 109)
	slice1 = append(slice1, 1987)
	fmt.Println(slice1)

	slice2 := array2[1:3] // é um ponteiro, indices 1, 2 apenas
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	fmt.Println("-------------------")
	fmt.Println("# Arrays Internos #")

	// o metodo make() aloca um espaço para algo, seja slice, etc
	// cria um array 'interno' de 15 posicoes e retorna um slice que vai pegar
	// as 10 primeiras posicoes do array

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 7) // quando o slice vai estourar, ele cria outro array com o dobro de tamanho
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	fmt.Println("************")
	slice4 := make([]float32, 5)
	//slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // tamanho
	fmt.Println(cap(slice4)) // capacidade
}
