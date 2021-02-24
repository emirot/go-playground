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

func secondFunc() {
	a := 11
	fmt.Println(a)
}

func returnWithNames() (first string, second string) {
	first = "first"
	//second = "second2"
	return
}

func main() {
	firstFunc()
	secondFunc()
	fmt.Println(firstRetFunc())
	fmt.Println(returnWithNames())
}
