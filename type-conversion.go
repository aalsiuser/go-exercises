package main

import (
	"fmt"
	"math"
)

func main() {
	var a,b = 4,5
	var c float64 = math.Sqrt(float64(a*b))
	fmt.Println(c)
}