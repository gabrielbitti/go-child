package main

import "fmt"

func main() {
	result := func(text string) string {
		// We should use specific for data type, example: %d for int, %s for string, etc.
		return fmt.Sprintf("Received: %s", text)
	}("Hello, World!!")

	fmt.Println(result)
}
