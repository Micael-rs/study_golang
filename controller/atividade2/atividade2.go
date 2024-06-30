package atividade2

import "fmt"

func Ponteiro() int {

	var valor1 int = 25

	a := &valor1
	*a = 10

	return valor1

}

func Vertex() {

	//
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

func Arrays() {

	//Declarando uma Array
	var a [5]int
	a[1] = 1
	fmt.Println(a)

	//manipulando uma array
	b := [5]int{}
	fmt.Println(b)

	//Criando uma fatia
	array := [5]int{1, 2, 3, 4, 5}
	var c []int = array[1:4]
	fmt.Println(c)

}
