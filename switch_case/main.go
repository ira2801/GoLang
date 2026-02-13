package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var y int32
	x := rand.Int31n(1000)

	switch {
	case x > 100:
		y = (x*x*x - 3) * (x*x + 3)
	default:
		y = x / (x*x - 4)
	}

	fmt.Printf("x: %d\n", x)
	fmt.Printf("Result: %d\n", y)
}
