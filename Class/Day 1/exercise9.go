package main

import (
	"fmt"
)

func main() {
	var coolNumbers = []int{24, 666, 69, 13, 9000, 10}
	var otherNumbers = []int{1, 2, 3, 4, 5, 6}
	
	var mehNumbers = append(coolNumbers, otherNumbers...)

	fmt.Println(coolNumbers)
	fmt.Println(otherNumbers)
	fmt.Println(mehNumbers)
}