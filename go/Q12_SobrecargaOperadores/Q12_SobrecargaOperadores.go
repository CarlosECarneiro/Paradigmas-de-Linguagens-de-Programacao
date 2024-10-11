package main

import (
	"fmt"
)

type Produto struct {
	nome  string
	preco float64
}

func (p Produto) Somar(outro Produto) float64 {
	return p.preco + outro.preco
}
func (p Produto) Info() string {
	return fmt.Sprintf("Produto: %s, Preço: R$ %.2f", p.nome, p.preco)
}

func main() {
	produto1 := Produto{nome: "Produto A", preco: 10.50}
	produto2 := Produto{nome: "Produto B", preco: 15.75}

	fmt.Println(produto1.Info())
	fmt.Println(produto2.Info())
	fmt.Printf("Soma dos produtos: %.2f", produto1.Somar(produto2)) 
}
