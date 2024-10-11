package main

import (
	"fmt"
)

type SaldoInsuficienteException struct {
	SaldoDisponivel float64
	ValorSaque     float64
}

func (e SaldoInsuficienteException) Error() string {
	return fmt.Sprintf("Saldo insuficiente: disponível R$ %.2f, tentativa de saque R$ %.2f", e.SaldoDisponivel, e.ValorSaque)
}

type ContaBancaria struct {
	saldo float64
}

func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > c.saldo {
		return SaldoInsuficienteException{
			SaldoDisponivel: c.saldo,
			ValorSaque:      valor,
		}
	}
	c.saldo -= valor
	return nil
}

func (c ContaBancaria) ExibirSaldo() float64 {
	return c.saldo
}

func main() {
	conta := ContaBancaria{saldo: 100.00}

	if err := conta.Sacar(150); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Saque realizado com sucesso. Saldo atual: R$ %.2f\n", conta.ExibirSaldo())
	}

	if err := conta.Sacar(50); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Saque realizado com sucesso. Saldo atual: R$ %.2f\n", conta.ExibirSaldo())
	}
}
