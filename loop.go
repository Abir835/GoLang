package main

import "fmt"

func main() {
	start := 1
	//things := []string{"maleta", "bolso"}
	//fmt.Println(start, len(things))
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i + start)
	//}
	//
	//for i := 0; i < len(things); i++ {
	//	fmt.Println(things[i])
	//}
	//
	//// range based loop
	//
	//for i := range things {
	//	fmt.Println(things[i])
	//}

	for start < 100 {
		start += start
		if start == 32 {
			goto lcolabel
		}
		fmt.Println(start)
	}

lcolabel:
	fmt.Println("lakhdlhfd")
}
