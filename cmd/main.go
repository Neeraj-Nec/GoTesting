package main

import (
	"fmt"
	"testPrime/pkg/losic"
)

func main() {
	n := 2
	_, result := losic.IsPrime(n)
	fmt.Println(result)
}
