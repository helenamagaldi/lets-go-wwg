package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(24)
	fmt.Println("My favorite number is", rand.Intn(10))
}