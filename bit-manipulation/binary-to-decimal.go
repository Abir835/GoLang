package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	num := "1101"
	ans := 0
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '1' {
			ans += int(math.Pow(2, float64(len(num)-1-i)))
		}
	}
	fmt.Println(ans)

	// if i have floating value

	binary := "1101.11"
	parts := strings.Split(binary, ".")
	integerParts := parts[0]
	facParts := ""
	if len(parts) > 1 {
		facParts = parts[1]
	}

	integerValue := 0
	for i := len(integerParts) - 1; i >= 0; i-- {
		if integerParts[i] == '1' {
			integerValue += int(math.Pow(2, float64(len(integerParts)-1-i)))
		}
	}

	facValue := 0.0
	for i := 0; i < len(facParts); i++ {
		if facParts[i] == '1' {
			facValue += math.Pow(2, float64(-(i + 1)))
		}
	}

	result := float64(integerValue) + facValue
	fmt.Println(result)
}
