package main

import (
	"fmt"
)

func main() {
	var days [7]string
	fmt.Println(days)
	
	days[0] = "Sunday bunday"
	days[1] = "Monday bad day"
	days[2] = "Tuesday Tacos"
	days[3] = "Wednesday Wings"
	days[4] = "Thursday 2 for exercuse"
	days[5] = "HAPPY HOUR FRIDAY"
	days[6] = "Saturday Night Fever"	
	fmt.Println(days)
	
	// also:
	days2 := [7]string{"Sunday bunday", "Monday bad day", "Taco Tuesday", "Wednesday Wings", "Thursday 2 for 1", "HAPPY HOUR FRIDAY", "Saturday Night Fever"}
	fmt.Println(days2)
}