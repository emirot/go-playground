package main

import (
	"fmt"
)

func firstFunc() {
	fmt.Println("First print")
}

func firstRetFunc() string {
	return "a string returnned"
}

func returnWithNames() (first string, second string) {
	first := "first"
	second := "second"
	return
}

func main() {
	firstFunc()
	fmt.Println(firstRetFunc())
	fmt.Println(returnWithNames())
}
