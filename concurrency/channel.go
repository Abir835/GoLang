package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	channel := make(chan string, 3)

	channel <- "hello world"
	channel <- "hello world2"
	channel <- "hello world3"

	//close(channel)
	//
	//for elem := range channel {
	//	fmt.Println(elem)
	//}
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	//go attack()
	//time.Sleep(time.Second)
}

//func attack() {
//	fmt.Println("attack")
//}
