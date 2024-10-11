package main

import (
	"fmt"
)

type Calculadora struct{}

func (c Calculadora) SomarDois(x, y float64) float64 {
	return x + y
}

func (c Calculadora) SomarTres(x, y, z float64) float64 {
	return x + y + z
}

func main() {
	calculadora := Calculadora{}

	resultado1 := calculadora.SomarDois(10, 20)
	fmt.Printf("A soma de 10 e 20 é: %.2f\n", resultado1)

	resultado2 := calculadora.SomarTres(10, 20, 30)
	fmt.Printf("A soma de 10, 20 e 30 é: %.2f\n", resultado2)
}
