package main

import (
	"fmt"
)

func firstFunc() {
	fmt.Println("First print")
}

func firstRetFunc() (string) {
	return "a string returnned"
}


func main() {
	firstFunc()
	fmt.Println(firstRetFunc())
}
