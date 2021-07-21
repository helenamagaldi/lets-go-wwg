// tipos compostos

package main

import (
	"fmt"
)

type Apartamento struct {
	numero              int
	nomeDaProprietaria  string
	possuiVagaDeGaragem bool
}

func main() {
	ap1 := Apartamento{
		numero:              13,
		nomeDaProprietaria:  "Bruxa do 71",
		possuiVagaDeGaragem: true,
	}
	fmt.Printf("%+v", ap1)

	fmt.Printf("O apartamento número %d tem como proprietária a %s. Ele tem vaga de garagem? %t", ap1.numero, ap1.nomeDaProprietaria, ap1.possuiVagaDeGaragem)
}
