package main

import (
	"fmt"
	"math"
)

func main() {
	q := 16.0
	r := 8.0

	s := math.Mod(q, r)
	fmt.Println(s)
}
