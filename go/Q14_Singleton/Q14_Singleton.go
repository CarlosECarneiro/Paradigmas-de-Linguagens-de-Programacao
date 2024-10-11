package main

import (
	"fmt"
	"sync"
)

type Configuracao struct {
}

var instancia *Configuracao
var once sync.Once

func GetConfiguracao() *Configuracao {
	once.Do(func() {
		instancia = &Configuracao{}
	})
	return instancia
}

func main() {
	config := GetConfiguracao()
	config2 := GetConfiguracao()
	fmt.Println( config == config2)
}
