package main

import "fmt"

func main() {
	idade := 69
	if idade >= 60 {
		fmt.Println("gold but gold")
		return
	}
	fmt.Println("young and dumb")
}