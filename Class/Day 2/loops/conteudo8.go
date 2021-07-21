package main

import (
	"fmt"
)

func main() {
	var meuSliceParaFor = []string{"for com contador", "for condicional", "for em algo", "E acabou a forezada!"}

	for indice, valor := range meuSliceParaFor {
		fmt.Printf("%d - %s\n", indice, valor)
	}

	for indice := range meuSliceParaFor {
		fmt.Printf("%d", indice)
	}
}