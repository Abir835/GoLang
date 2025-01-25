package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	var things = []string{"maleta", "ropa", "reloj"}
	fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)

	things = append(things[1 : len(things)-1])
	fmt.Println(things)

	hero := make([]string, len(things))
	fmt.Println(hero)

	myInt := []int{1, 2, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(myInt)))
	fmt.Println(myInt)

	// slice is simmer to array
	slice := []int{1, 2, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	fmt.Println(slice)
	slice = append(slice, 4)
	fmt.Println(slice)
}
