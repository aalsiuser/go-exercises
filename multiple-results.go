package main

import (
	"fmt"
)

func swap(a, b string) (string, string, string) {
	return b, a, a+b
}

func main(){
	a,b,c := swap("hello", "vamshi")
	fmt.Println(a,b,c)
}