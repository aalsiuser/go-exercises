package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x")
	case "linux":
		fmt.Println("OS linux")
	default:
		fmt.Printf("%s. \n", os)
	}
}