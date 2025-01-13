package main

import (
	"fmt"
)

//func fibonacci(c, quit chan int) {
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}
//
//func main() {
//	c := make(chan int)
//	quit := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c)
//		}
//		quit <- 0
//	}()
//
//	fibonacci(c, quit)
//}

func main() {
	msgChan := make(chan string)
	go func() {
		//time.Sleep(1 * time.Second)
		msgChan <- "Hello"
		msgChan <- "World"
	}()

	msg1 := <-msgChan
	msg2 := <-msgChan

	fmt.Println(msg1, msg2)

}
