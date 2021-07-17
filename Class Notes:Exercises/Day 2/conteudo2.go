package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

type Aula struct {
	Descricao string
	Docente   Pessoa
	Discentes []Pessoa
}

func main() {

	eliana := Pessoa{Nome: "Eliana", Idade: 47}
	joao := Pessoa{Nome: "João", Idade: 19}
	zeca := Pessoa{Nome: "Zeca", Idade: 34}
	aline := Pessoa{Nome: "Aline", Idade: 42}

	aula := Aula{
		Descricao: "Aula de Go",
		Docente:   eliana,
		Discentes: []Pessoa{
			joao,
			zeca,
			aline,
		},
	}

	fmt.Printf("%+v", aula)
}