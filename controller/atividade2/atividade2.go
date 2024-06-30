package atividade2

import "fmt"

func Ponteiro() int {

	var valor1 int = 25

	a := &valor1
	*a = 10

	return valor1

}

func Vertex() {

	type Vertex struct {
		Valor1 int
		Valor2 int
	}

	v := Vertex{1, 2}
	v.Valor1 = 4

	fmt.Println(v.Valor2)

	p := &v
	p.Valor2 = 8

	fmt.Println(v.Valor2)

}
