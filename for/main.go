package main

import "fmt"

func main() {
	maxNumber := 0

	for i := 0; i <= 1_000_000_000; i++ {
		maxNumber = i
	}
	fmt.Println("Max number: ", maxNumber)

	init := 1
	// FOR IS GO's WHILE
	for init < 100 {
		init += init
	}
	fmt.Println(init)

	for i := range 5 {
		fmt.Println(i)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := range 100 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
