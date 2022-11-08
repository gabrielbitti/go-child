package main

import "fmt"

func main() {
	// serve para guardar dados
	// chave valor sempre do mesmo tipo
	// nao mutavel

	// para criar um map é necessario passar o tipo da chave entre colchetes
	//
	fmt.Println("# Maps")
	user := map[string]string{
		"first_name": "Jaws",
		"last_name":  "Loys",
	}

	fmt.Println(user)
	fmt.Println(user["first_name"])
	fmt.Println(user["last_name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "Joao",
			"last":  "Batista",
		},
		"course": {
			"name": "Engenharia",
		},
	}

	fmt.Println("*****************")
	fmt.Println(user2)
	fmt.Println(user2["name"]["first"])
	fmt.Println(user2["name"]["last"])
	fmt.Println(user2["course"]["name"])
	delete(user2, "name")
	fmt.Println(user2)
}
