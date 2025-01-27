package main

import "fmt"

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {

	arr := []bool{}
	pre := [][]int{{1, 0}}
	query := [][]int{{1, 0}, {0, 1}}

	// If prerequisites are empty, append `false` for each query
	if len(pre) == 0 {
		for range query {
			arr = append(arr, false)
		}
	} else {
		for _, i := range query {
			found := false
			for _, j := range pre {
				if IntArrayEquals(i, j) {
					found = true
					break
				}
			}
			arr = append(arr, found)
		}
	}

	fmt.Println(arr)
}
