package main

import (
	"fmt"
)

type Motor struct {
	 potencia int
	 tipo string
}

type Carro struct {
	marca string 
	modelo string 
	motor  Motor 
}
func (c Carro) ligar(){
	fmt.Println("O carro está pronto para rodar!")
}

func main() {
	motor := Motor{potencia: 150, tipo: "Gasolina"}
	carro := Carro{marca: "Toyota", modelo: "Corolla", motor: motor}
	carro.ligar()
}
