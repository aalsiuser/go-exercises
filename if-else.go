package main

import (
	"fmt"
	"math"
)

func pow(x,y,z float64) float64 {
	if v := math.Pow(x,y); v < z {
		return v
	} else{
		fmt.Println("%g >= %g\n", v, z)
	}

	return z
}
func main() {
	fmt.Println(pow(3, 3, 20))
}