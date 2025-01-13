package main

import "fmt"

func main() {
	var p *int

	a := 20

	p = &a
	fmt.Println(*p)

}
