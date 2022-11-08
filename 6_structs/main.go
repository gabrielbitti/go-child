// é o mais proximo de uma classe. no go nao existem classes
// é uma coleção de campos, onde cada campo tem um tipo

package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint16
}

func main() {
	fmt.Println("# Structs #")

	var user1 user
	user1.age = 1
	user1.name = "Eistrockensz"

	user2Address := address{"Street floor, NY, USA", 25}
	user2 := user{"Carlito F.", 20, user2Address}

	user3 := user{name: "Curizotinio"}

	fmt.Println("## User 1: ", user1)
	fmt.Println("## User 1 Address: ", user1.address)

	fmt.Println("## User 2: ", user2)
	fmt.Println("## User 2 Address: ", user2.address)

	fmt.Println("## User 3: ", user3)
	fmt.Println("## User 3 Address: ", user3.address)
}
