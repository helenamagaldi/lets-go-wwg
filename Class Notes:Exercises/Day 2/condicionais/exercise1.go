package main

import "fmt"

func main() {
	nascimento := 1998 
	hoje := 2021
	if hoje - nascimento >= 16 {
		fmt.Println("OLD BUT GOLD GO VOTE")
		return
	}
	
	fmt.Println("young and dumb, can't vote")
}