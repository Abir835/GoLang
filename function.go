package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(add(3, 2))
	myResult, name, length := addValue(1, 2, 3, 4, 5)
	fmt.Println("My Result", myResult, name, length)
	fmt.Println("Int to String", convertIntToString(1))
}

func add(x int, y int) int {
	fmt.Printf("x: %d\n", x+y)
	return multiply(x, y)
}

func multiply(x, y int) int {
	return x * y
}

func convertIntToString(a int) string {
	return strconv.Itoa(a)
}

func addValue(values ...int) (int, string, int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	name := "ABIR HASAN"
	length := len(name)
	return sum, name, length
}
