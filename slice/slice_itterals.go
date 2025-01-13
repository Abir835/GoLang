package main

import "fmt"

func main() {
	s := []struct {
		name string
		age  int
	}{
		{"Abir Hasan", 21},
		{"Abir Hasan", 21},
		{"Abir Hasan", 21},
		{"Abir Hasan", 21},
		{"Abir Hasan", 21},
		{"Abir Hasan", 21},
		{"Abir Hasan", 21},
		{"Abir Hasan", 21},
	}

	for _, v := range s {
		fmt.Println(v)
	}

	//fmt.Println(s)
}
