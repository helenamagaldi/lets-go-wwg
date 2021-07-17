package main

import (
	"fmt"
)

type Pessoa struct {
	nome         string
	idade        int
	corPreferida string
}

func main() {
	godOfWar := Pessoa{
		nome:         "Kratos",
		idade:        1234,
		corPreferida: "blood of my enemies",
	}

	hellzBellz := Pessoa{
		nome:         "Helena Magaldi",
		idade:        33,
		corPreferida: "gold",
	}

	darkElf := Pessoa{
		nome:         "Muria",
		idade:        666,
		corPreferida: "toxic Green",
	}

	fmt.Printf("O primeiro personagem é o %s, ele tem %d anos e a sua cor preferida é %s\n", godOfWar.nome, godOfWar.idade, godOfWar.corPreferida)
	fmt.Printf("A segunda personagem é a %s, ela tem %d anos e a sua cor preferida é %s\n", hellzBellz.nome, hellzBellz.idade, hellzBellz.corPreferida)
	fmt.Printf("A terceira personagem é a %s, ela tem %d anos e a sua cor preferida é %s\n", darkElf.nome, darkElf.idade, darkElf.corPreferida)
}
