package main

import (
	"fmt"
)

type Cat struct {
	name string
	mom  string
	food string
}

func (g Cat) Hello() {
	fmt.Printf("=^._.^= ∫ Meow, my name is %s, mommy is called %s and my favorite food is %s\n", g.name, g.mom, g.food)

}

type Dog struct {
name string
mom string
toy string
}

func (d Dog) Hello() {
	fmt.Printf("V●ω●V. ∫ Bark, my name is %s, mommy is called %s and my favorite food is %s\n", d.name, d.mom, d.toy)
}

func main() {
	cat1 := Cat{
		name: "Fred",
		mom:  "Helena",
		food: "lamb",
	}
	cat2 := Cat{
		name: "Francisco",
		mom:  "Helena R.",
		food: "salmon",
	}

	dog1 := Dog{
	name: "Chicória",
	mom: "Amanda",
	toy: "everything",
	} 
	
	
	cat1.Hello()
	cat2.Hello()
	
	dog1.Hello()
}
