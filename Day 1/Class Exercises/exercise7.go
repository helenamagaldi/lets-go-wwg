package main

import (
	"fmt"
)

func main() {
	elements := [5]string{"zero", "one", "two", "three", "four"}
	fmt.Println(elements)
	
	// especificador de formato e tipo
	fmt.Printf("o tipo é: %T\n", elements)
}