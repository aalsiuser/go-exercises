package main

import (
	"fmt"
	"math"
)

func pow(x,y,z float64) float64 {
	if v := math.Pow(x,y); x < z {
		return v
	}

	return z
}
func main() {
	fmt.Println(pow(3,2, 10))
}