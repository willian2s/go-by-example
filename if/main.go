package main

import "fmt"

func isOdd(num int) string {
	if num%2 == 0 {
		return fmt.Sprintf("%d is Even", num)
	}
	return fmt.Sprintf("%d is Odd", num)
}

func isDivisible(num, div int) string {
	if num%div == 0 {
		return fmt.Sprintf("%d is divisible by %d.", num, div)
	}
	return fmt.Sprintf("%d is not divisible by %d.", num, div)
}

func main() {
	fmt.Println(isOdd(7))
	fmt.Println(isOdd(8))
	fmt.Println(isDivisible(8, 4))
	fmt.Println(isDivisible(8, 3))

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
