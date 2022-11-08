package main

import (
	"aula1/auxx"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("From main.")
	auxx.Write()
	err := checkmail.ValidateFormat("gabrielbitti0@gmail.com")
	fmt.Println(err)
}
