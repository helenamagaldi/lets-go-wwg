// Exercicio 1

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var quilometros int8
// 	quilometros = 150

// 	fmt.Println(quilometros)
// }

// 1.1
// O programa não compila pois o valor da variavel quilometros excede o valor determinado acima, na declaração do tipo da variável (no caso, int8, que tem um valor máximo de 126)

// 1.2
// O erro nos indica que o valor 150 não é aceitável pois supera o valor máximo de int8, que é 126

// 1.3
// CORREÇÃO DO CÓDIGO ABAIXO:

package main

import (
	"fmt"
)

func main() {
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}