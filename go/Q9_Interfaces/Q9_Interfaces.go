package main

import (
	"fmt"
)

type Imprimível interface {
	Imprimir()
}

type Relatorio struct {
	conteudo string
}

func (r Relatorio) Imprimir() {
	fmt.Printf("Imprimindo relatório: %s\n", r.conteudo)
}

type Contrato struct {
	conteudo string
}

func (c Contrato) Imprimir() {
	fmt.Printf("Imprimindo contrato: %s\n", c.conteudo)
}


func main() {
	relatorio := Relatorio{conteudo: "Relatório"}

	contrato := Contrato{conteudo: "Contrato"}

	relatorio.Imprimir()
	contrato.Imprimir()
}
