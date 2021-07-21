package main

import "fmt"

func main() {
	hora := 0

	for hora < 3 {
		for minuto := 0; minuto < 60; minuto++ {
			for segundos := 0; segundos < 60; segundos++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundos)
			}
		}

		hora++
	}

}
