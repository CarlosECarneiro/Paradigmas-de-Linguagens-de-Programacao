package main

import (
	"fmt"
)

type Professor struct {
	nome      string
	escolas   []*Escola 
}

type Escola struct {
	nome      string
	professores []*Professor 
}

func (e *Escola) AdicionarProfessor(p *Professor) {
	e.professores = append(e.professores, p)
	p.escolas = append(p.escolas, e)
}

func (e Escola) Info() string {
	return fmt.Sprintf("Escola: %s, Professores: %d", e.nome, len(e.professores))
}

func (p Professor) Info() string {
	return fmt.Sprintf("Professor: %s, Escolas: %d", p.nome, len(p.escolas))
}

func main() {
	prof1 := &Professor{nome: "João"}
	prof2 := &Professor{nome: "Maria"}

	escola1 := Escola{nome: "Escola A"}
	escola2 := Escola{nome: "Escola B"}

	escola1.AdicionarProfessor(prof1)
	escola1.AdicionarProfessor(prof2)
	escola2.AdicionarProfessor(prof1)

	fmt.Println(escola1.Info())
	fmt.Println(escola2.Info())

	fmt.Println(prof1.Info())
	fmt.Println(prof2.Info())
}
