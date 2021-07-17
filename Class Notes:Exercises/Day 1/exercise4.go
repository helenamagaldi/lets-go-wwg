package main

import (
	"fmt"
)

func main() {
	var good, bad, late, cringe, millenial = true, false, false, true, true
	
	fmt.Printf("type is %T, it's value is %v\n", good, good)
	fmt.Printf("type is %T, it's value is %v\n", bad, bad)
	fmt.Printf("type is %T, it's value is %v\n", late, late)
	fmt.Printf("type is %T, it's value is %v\n", cringe, cringe)
	fmt.Printf("type is %T, it's value is %v\n", millenial, millenial)
	
	years := 2021 > 2020
	time := 1 <= 2
	scale := 9000 == 10
	
	fmt.Printf("type is %T, it's value is %v\n", years, years)
	fmt.Printf("type is %T, it's value is %v\n", time, time)
	fmt.Printf("type is %T, it's value is %v\n", scale, scale)		
}