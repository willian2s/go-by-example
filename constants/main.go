package main

import "fmt"

func main() {
	const name = "Caio"
	const age = 11
	const mother = "Mariana"
	const father = "Willian"
	const sister = "Helena"

	const canDrive = age >= 18

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Mother: %s\n", mother)
	fmt.Printf("Father: %s\n", father)
	fmt.Printf("Sister: %s\n", sister)
	fmt.Printf("Can drive? %t\n", canDrive)
}
