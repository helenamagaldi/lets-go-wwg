package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	birthday := time.Date(1989, 1, 2, 0, 0, 0, 0, time.UTC)
	today := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	//today := time.Now()

	fmt.Println(math.Floor(today.Sub(birthday).Hours() / 24 / 365))
}
