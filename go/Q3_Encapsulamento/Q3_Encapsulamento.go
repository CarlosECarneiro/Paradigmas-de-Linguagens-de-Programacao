package main

import "fmt"

type ContaBancaria struct {
	titular string
	saldo float64
}
func (c *ContaBancaria) Depositar(dinheiro float64) {
	c.saldo+=dinheiro
}
func (c *ContaBancaria) Sacar(dinheiro float64) {
	c.saldo-=dinheiro
}
func main(){
	conta := ContaBancaria{titular: "João",saldo: 100}
	conta.Depositar(200)
	conta.Sacar(50)
	fmt.Printf("%f",conta.saldo)
}