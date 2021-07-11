package main

import (
	"fmt"
)

func main() {
	var itsSunnyOutside bool
	itsSunnyOutside = true
	
	var itsRaining bool = false
	
	var itsCold = true
	

	fmt.Printf("type is %T, it's value is %v\n", itsSunnyOutside, itsSunnyOutside)
	fmt.Printf("type is %T, it's value is %v\n", itsRaining, itsRaining)
	fmt.Printf("type is %T, it's value is %v\n", itsCold, itsCold)
}

