package main

import (
	"fmt"
)

func main() {
	var groceries = []string{"basil", "olive oil", "rice & beans"}
	fmt.Println(groceries)

	var groceriesSis = []string{"shampoo", "bread"}

	groceries2 := append(groceries, "chocolate", "pasta")
	fmt.Println(groceries2)	
	
	var groceriesTotal = append(groceries, groceriesSis...)
	fmt.Println(groceriesTotal)
}