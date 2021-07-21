package main

import (
	"fmt"
)

func main() {
	var fatia = []string{"zero", "um", "dois", "tres", "quatro"}
	for index, value := range fatia {
		fmt.Println(index, value)
	}
}