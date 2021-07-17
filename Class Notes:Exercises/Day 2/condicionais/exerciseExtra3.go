package main

import "fmt"

func main() {
	hora := 13
	
	switch {
	case hora >= 6 && hora < 12:
	fmt.Println("é manhã")
	case hora >= 12 && hora < 18:
	fmt.Println("é tarde")
	case hora >= 18 && hora < 24:
	fmt.Println("é noite")
	case hora >= 0 && hora < 6:
	fmt.Println("é madrugada")
	}
}
