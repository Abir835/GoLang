package main

import (
	"fmt"
	"math"
)

var pow = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	for i, v := range pow {
		b := math.Pow(2, float64(v))
		fmt.Printf("2**%d = %d\n", i, int(b))
	}
}
