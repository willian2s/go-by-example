package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("macOS")
	default:
		fmt.Println(os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("To far way")
	}

	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning")
	case time.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
