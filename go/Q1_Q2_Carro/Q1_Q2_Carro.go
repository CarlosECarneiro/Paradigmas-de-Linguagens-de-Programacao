// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Carro struct {
	marca, modelo string
	ano int
	velocidade int
}
func NewCarro (marca string, modelo string, ano int) Carro{
	carro := Carro{}
	carro.marca = marca
	carro.modelo = modelo
	carro.ano = ano
	carro.velocidade = 0
	return carro
}
func (c *Carro) Acelerar() {
	c.velocidade+=10
}
func (c *Carro) Frear() {
	c.velocidade-=10
}
func (c Carro) ExibirVelocidade(){
	fmt.Printf("Velocidade atual: %d\n", c.velocidade)
}
func main() {
	carro1 := NewCarro("Fiat", "Uno", 1998)
	carro2 := NewCarro("Volkswagen","Golf", 2014)
	carro3 := NewCarro("Honda", "Civic", 2023)
	fmt.Printf("%s %s %d\n",carro1.marca,carro1.modelo,carro1.ano)
	fmt.Printf("%s %s %d\n",carro2.marca,carro2.modelo,carro2.ano)
	fmt.Printf("%s %s %d\n",carro3.marca,carro3.modelo,carro3.ano)
	carro1.ExibirVelocidade()
	carro1.Acelerar()
	carro1.ExibirVelocidade()
	carro1.Frear()
	carro1.ExibirVelocidade()
}
