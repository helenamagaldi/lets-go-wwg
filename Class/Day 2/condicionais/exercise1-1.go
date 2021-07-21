package main

import "fmt"

func main() {
	nascimento := 1998 
	hoje := 2021
	age := hoje - nascimento
	
	if age >= 16 {
		fmt.Printf("OLD BUT GOLD GO VOTE, this person has %d years old\n", age)
		return
	} else {
		fmt.Printf("young and dumb, can't vote, they have %d years old\n", age)	
	}
}