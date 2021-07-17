package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Please, enter your birthday year: ")
	var birthday int
	fmt.Scanln(&birthday)

	today := time.Now().Year()
	var todayInt = int(today)

	yearsOld := todayInt - birthday
	fmt.Printf("You're %v years old", yearsOld)

}