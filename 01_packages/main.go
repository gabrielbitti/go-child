package main

import (
	"01_packages/auxx"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("# Packages")
	fmt.Println("From main.")
	auxx.Write()

	mail1 := "gabrielbitti0@gmail.com"
	mail2 := "s@@@a2a.co0m"

	err := checkmail.ValidateFormat(mail1)
	if err != nil {
		fmt.Println(mail1, "is not a valid mail address.")
	} else {
		fmt.Println(mail1, "is a valid mail address.")
	}

	err = checkmail.ValidateFormat(mail2)
	if err != nil {
		fmt.Println(mail2, "is not a valid mail address.")
	} else {
		fmt.Println(mail2, "is a valid mail address.")
	}
}
