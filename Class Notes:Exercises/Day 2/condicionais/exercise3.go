package main

import "fmt"

func main() {
idade := 47

if idade >= 60 {
fmt.Printf("a pessoa tem idade %d e é idosa\n", idade)
}else if idade >= 18 && idade < 60 {
fmt.Printf("a pessoa tem idade %d e é maior de idade\n", idade)
} else if idade < 0 {
fmt.Println("o valor idade é negativo e, portante, inválido")
} else {
fmt.Printf("a pessoa tem idade %d e é menor de idade\n", idade)}

}
