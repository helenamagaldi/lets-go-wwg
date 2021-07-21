package main

import (
	"fmt"
)

func main() {
	list := []string{"chocolate", "cake", "lollipop", "milkshake", "diabetes"}
	for index, value := range list {
		fmt.Printf("%d) %s\n",index+1, value)
	}
}