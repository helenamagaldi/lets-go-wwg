package main

import (
	"fmt"
)

func main() {
	var name, color string = "hells bellz", "gold"
	
	fmt.Printf("My name is %s and my favorite color is %s\n", name, color)

	color = "black"
	
	fmt.Printf("My name is %s and my favorite color is now %s\n", name, color)
}
