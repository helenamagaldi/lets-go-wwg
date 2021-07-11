{TODO: formatting}
# https://womenwhogocwb.gitbook.io/letsgo/tipos-de-dados-basicos/exercicios-extras

## Exercício #01
A aplicação abaixo não compila.

package main

import (
	"fmt"
)

func main() {
	var quilometros int8
	quilometros = 150

	fmt.Println(quilometros)
}


1) Descubra por que não compila
2) Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso código. O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
3) Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela
Tamanhos e valores do tipo inteiro com sinal
Tipo, Tamanho, espectro de valores:
int8, 8bits, -128 ~ 127
int16, 16bits, -2^15 ~ 2^15-1
int32, 32bits, -2^31 ~ 2^31-1
int64, 64bits, -2^63 ~ 2^63-1

## Exercício #02
1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
2) Como podemos pegar a informação do ano diretamente do console?


Dicas: 
O método time.Now() retorna a data atual (para ler mais sobre ele na documentação de Go, clique aqui)
Para ler do console, você pode utilizar o método fmt.Scan() (para ler mais sobre ele na documentação de Go, clique aqui)