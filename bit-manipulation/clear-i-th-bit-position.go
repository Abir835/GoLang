package main

import "fmt"

func main() {
	n := 13
	i := 2

	res := n & ^(1 << i)
	fmt.Println(res)

}
