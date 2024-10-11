package main

import (
	"fmt"
)

type Empregado struct {
	nome   string
	salario float64
}

func (e Empregado) Info() string {
	return fmt.Sprintf("Nome: %s, Salário: R$%.2f", e.nome, e.salario)
}

type Empresa struct {
	nome      string
	empregados []*Empregado 
}

func (emp *Empresa) AdicionarEmpregado(e *Empregado) {
	emp.empregados = append(emp.empregados, e)
}

func (emp Empresa) Info() {
	fmt.Printf("Empresa: %s\n", emp.nome)
	fmt.Println("Empregados:")
	for _, e := range emp.empregados {
		fmt.Println(e.Info())
	}
}

func main() {
	empresa := Empresa{nome: "Empresa abc"}

	emp1 := &Empregado{nome: "João", salario: 1600.00}
	emp2 := &Empregado{nome: "Maria", salario: 5000.00}

	empresa.AdicionarEmpregado(emp1)
	empresa.AdicionarEmpregado(emp2)

	empresa.Info()
}
