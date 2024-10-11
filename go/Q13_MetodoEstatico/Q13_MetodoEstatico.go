package main

import (
	"fmt"
)

type Matematica struct{}

func (m Matematica) Fatorial(n int) int {
	if n < 0 {
		return -1
	}
	resultado := 1
	for i := n; i > 0; i-- {
		resultado *= i
	}
	return resultado
}

func main() {
	matematica := Matematica{}

	fmt.Printf("%d",matematica.Fatorial(5))
}
