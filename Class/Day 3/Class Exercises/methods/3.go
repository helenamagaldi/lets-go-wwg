package main

import (
	"fmt"
)

// Considere um cenário em que você precise regularmente trabalhar com slices de inteiros
//e extrair a soma e média dos números dessa lista. Usando o que você aprendeu sobre métodos,
//qual seria sua solução para esse problema em Go?

type conjunto []int

func (c conjunto) Some() int {
	soma := 0
	for _, operando := range c {
		soma += operando
	}
	return soma
}

func (c conjunto) CalculeMedia() float64 {
	soma := float64(c.Some())
	quantidade := float64(len(c))
	return soma / quantidade
}

func main() {
	conjunto1 := conjunto{3, 7, 27, 40}
	conjunto2 := conjunto{5, 10, 25, 3, 7}

	var conjuntos = []conjunto{conjunto1, conjunto2}

	for _, valor := range conjuntos {
		soma := valor.Some()
		media := valor.CalculeMedia()
		fmt.Printf("nosso conjunto tem os elementos: %v\n\tsoma: %d\n\tmédia: %f\n", valor, soma, media)
	}
}
