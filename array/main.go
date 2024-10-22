package main

import "fmt"

func main() {
	// create empty array with lenght 5
	var a [5]int
	fmt.Println("emp:", a)

	// set a value in index 4
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// get array lenght
	fmt.Println("len:", len(a))

	// declare a new array with value
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// omit number of elements
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// when pass a index, the elements in between will be zeroed
	b = [...]int{100, 3: 200, 300}
	fmt.Println("dcl:", b)

	// declare a 2d array
	// array are one-dimensional, but you can compose types
	// to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)

	// initialize a 2D array
	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D:", twoD)
}
