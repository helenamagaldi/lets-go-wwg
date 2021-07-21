package main

import (
	"fmt"
)

func main() {
	var nome = "Alaska Thunderfuck"
	idade := 24
	fmt.Println(nome, idade)
	fmt.Printf("O nome é %s e a idade é %d\n", nome, idade)
	fmt.Printf("O tipo de nome é %T e o tipo de idade é %T\n\n", nome, idade)
	
	var precoDoGlitter, pesoDoGlitter = 2.800, 1.480
	fmt.Println("total: ", precoDoGlitter * pesoDoGlitter)
	fmt.Printf("%T, %T\n\n", precoDoGlitter, pesoDoGlitter)
	
	sabadaoPagodao := true
	sabadaoJulho := true
	temLetsGo := (sabadaoPagodao && sabadaoJulho)
	fmt.Println(temLetsGo, "\n")

	var nomes = [3]string{"hellz", "bellz", "hellzbellz"}
	fmt.Println(nomes, "\n")
	
	idades := make([]int,3)
	idades[0] = 24
	fmt.Println(idades)
	idades[1] = 42
	idades[2] = 69
	fmt.Println(idades)
	idades = append(idades, 666)
	fmt.Println(idades, "\n")
	
	var mapa = make(map[string]int)
	mapa["X"] = 10
	fmt.Println(mapa)
	fmt.Println(mapa["X"], "\n")
	
	cores := map[string]string{
	"cinza": "#cccccc",
	"roxo": "#a378f8",
	"branco": "#ffffff"
	}
	cinza := cores["cinza"]
	fmt.Println(cinza)
	azul, existe := cores["azul: ]
	fmt.Println(azul, existe)
	
	fmt.Println(cores)
	delete(cores, "branco")
	fmt.Println(cores)

	
}
