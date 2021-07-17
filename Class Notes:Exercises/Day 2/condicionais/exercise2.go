package main

import "fmt"

func main() {
	variable  := 24

	if variable > 0 {
		fmt.Printf("Variable with %d number: positive\n", variable) 
		return
	} else {
		fmt.Printf("Variable with %d number: negative\n", variable)
	}
}
