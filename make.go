package main

import "fmt"

func main() {
	a := make([]int, 0, 5)
	b := append(a, 1)
	fmt.Println(b)
}
