package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var myString string
	//fmt.Scanln(&myString)
	//fmt.Println(myString)

	//var name string = "Abir Hasan"
	//var val, _ = fmt.Println(name)
	//fmt.Println(val)
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter Username: ")
	//username, _ := reader.ReadString('\n')
	//fmt.Print("Enter Password: ")
	//password, _ := reader.ReadString('\n')
	//fmt.Println(username)
	//fmt.Println(password)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Rating: ")
	myRating, _ := reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	fmt.Println("Your Rating is: ", myNumRating+2)
}
