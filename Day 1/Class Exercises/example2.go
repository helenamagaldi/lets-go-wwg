package main

import (
	"fmt"
)

func main() {
	ensolarado := true
	disposta := true
	sair := false
	
	beach := ensolarado && disposta && sair
	
	fmt.Println(beach)
	fmt.Printf("%T", beach)
}