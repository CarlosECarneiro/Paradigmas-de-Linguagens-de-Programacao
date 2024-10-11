package main

import (
	"fmt"
)

type Animal struct {
	nome    string
	especie string
}



func (a Animal) Som() string {
	return ""
}

type Cachorro struct {
	Animal 
}

func (c Cachorro) Som() string {
	return "Au Au"
}

type Gato struct {
	Animal
}

func (g Gato) Som() string {
	return "Miau"
}


func main() {
	cachorro := Cachorro{Animal: Animal{nome:"Bicho",especie: "Cachorro"},}

	gato := Gato{Animal: Animal{nome:"Bichano",especie: "Gato"},}

	fmt.Println(cachorro.Som())
	fmt.Println(gato.Som())
}
