// nao é bem herança
package main

import "fmt"

type people struct {
	name     string
	lastName string
	age      uint8
}

type student struct {
	people  // isso é "herança" -- não é necessário colocar o tipo
	course  string
	college string
}

func main() {
	people1 := people{name: "Gabriel", lastName: "Bitti", age: 24}
	student1 := student{course: "Computer Science", college: "UDES", people: people1}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.lastName)
}
