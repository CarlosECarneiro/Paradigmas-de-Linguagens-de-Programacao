package main

import (
	"fmt"
)

type Animal interface {
	Som() string
}

type AnimalBase struct {
	nome    string
	especie string
}

func (a AnimalBase) Som() string {
	return ""
}

type Cachorro struct {
	AnimalBase 
}

func (c Cachorro) Som() string {
	return "Au Au"
}

type Gato struct {
	AnimalBase
}

func (g Gato) Som() string {
	return "Miau"
}

func emitirSons(animais []Animal) {
	for _, a := range animais {
		fmt.Println(a.Som())
	}
}

func main() {
	cachorro := Cachorro{AnimalBase: AnimalBase{nome: "Bicho", especie: "Cachorro"}}
	gato := Gato{AnimalBase: AnimalBase{nome: "Bichano", especie: "Gato"}}

	fmt.Println(cachorro.Som())
	fmt.Println(gato.Som())

	animais := []Animal{cachorro, gato}
	emitirSons(animais)
}
