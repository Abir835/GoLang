package main

import (
	"fmt"
)

func minFlagstones(n, m, a int64) int64 {
	return ((n + a - 1) / a) * ((m + a - 1) / a)
}

func main() {
	var n, m, a int64
	fmt.Scan(&n, &m, &a)
	fmt.Println(minFlagstones(n, m, a))
}
