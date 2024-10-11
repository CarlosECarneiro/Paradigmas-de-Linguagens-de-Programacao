package main

import (
	"fmt"
)

type Funcionario interface {
	CalcularSalario() float64
}

type FuncionarioHorista struct {
	nome         string
	horasTrabalhadas int
	valorHora    float64
}

func (f FuncionarioHorista) CalcularSalario() float64 {
	return float64(f.horasTrabalhadas) * f.valorHora
}

type FuncionarioAssalariado struct {
	nome  string
	salario float64
}

func (f FuncionarioAssalariado) CalcularSalario() float64 {
	return f.salario
}

func exibirSalario(funcionario Funcionario) {
	fmt.Printf("Salário: R$ %.2f\n", funcionario.CalcularSalario())
}

func main() {
	funcionario1 := FuncionarioHorista{nome:"Zé", horasTrabalhadas: 150, valorHora: 20.00}
	funcionario2 := FuncionarioAssalariado{nome:  "João",salario: 2000.00}

	fmt.Printf("Funcionário: %s\n", funcionario1.nome)
	exibirSalario(funcionario1)

	fmt.Printf("Funcionário: %s\n", funcionario2.nome)
	exibirSalario(funcionario2)
}
