package main

import "fmt"

func main() {
	// count set bit use right shift
	n := 13
	count := 0
	for i := 0; i < 4; i++ {
		set := (n >> i) & 1
		if set != 0 {
			count++
		}
	}
	fmt.Println(count)

	// now count set bits use left shift
	countSetBit := 0
	m := 13

	for i := 0; i < 4; i++ {
		x := m & (1 << i)
		if x != 0 {
			countSetBit++
		}
	}
	fmt.Println(countSetBit)

	// another way i cal solve
	z := 13
	c := 0

	for z > 0 {
		c += z & 1
		fmt.Println(c)
		z /= 2
	}

	fmt.Println(c)
}
