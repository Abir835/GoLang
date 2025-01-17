package main

import "fmt"

func main() {

	n := 13
	i := 2

	// left shift use
	val := n & (1 << i)
	if val != 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
