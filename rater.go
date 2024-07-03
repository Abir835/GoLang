package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//fmt.Println("Hello World")
	var name string
	var userRating string

	// front end
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	name, _ = reader.ReadString('\n')

	// take rating from user convert it to int
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter rate our Dosa center between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	// backend

	fmt.Printf("\n\nHello %v, \n Thanks for rating our dosa center with %v"+
		" star rating. \n\n Your rating was recorded in "+
		"our system at %v \n", name, myNumRating, time.Now().Format(time.Kitchen))

	if myNumRating == 5 {
		fmt.Println("Bonus for team for 5 star service")
	} else if myNumRating == 4 || myNumRating == 3 {
		fmt.Println("We are always improving")
	} else {
		fmt.Println("Need Serious work on our side")
	}
}
