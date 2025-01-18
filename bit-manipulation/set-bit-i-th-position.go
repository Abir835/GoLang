package main

import "fmt"

func main() {
	n := 13
	i := 2

	res := n | (i << 1)
	fmt.Println(res)
}
