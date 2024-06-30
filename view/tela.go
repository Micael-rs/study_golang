package view

import (
	"fmt"
	"study_golang/controller/atividade"
	"study_golang/controller/atividade2"
)

var valor1, valor2 int
var operacao = 1
var nome, email string

func TelaMenu() {

	for operacao != 0 {

		fmt.Println("Menu")
		fmt.Println("1-soma")
		fmt.Println("2-Subtração")
		fmt.Println("3-Multiplicação")
		fmt.Println("4-Divisão")
		fmt.Println("5-Atividade 2")
		fmt.Println("6-Cadastro")
		fmt.Println("0-Sair")
		fmt.Scanf("%d", &operacao)
		Operacao(operacao)
	}
}

func Operacao(int) {

	switch operacao {

	case 1:

		fmt.Println("Insira o primeiro valor:")
		fmt.Scanf("%d", &valor1)
		fmt.Println("Insira o segundo valor:")
		fmt.Scanf("%d", &valor2)

		fmt.Println("Resultado: ", atividade.Soma(valor1, valor2))

	case 2:

		fmt.Println("Insira o primeiro valor:")
		fmt.Scanf("%d", &valor1)
		fmt.Println("Insira o segundo valor:")
		fmt.Scanf("%d", &valor2)

		fmt.Println("Resultado: ", atividade.Subtracao(valor1, valor2))

	case 3:

		fmt.Println("Insira o primeiro valor:")
		fmt.Scanf("%d", &valor1)
		fmt.Println("Insira o segundo valor:")
		fmt.Scanf("%d", &valor2)

		fmt.Println("Resultado: ", atividade.Multiplicacao(valor1, valor2))

	case 4:

		fmt.Println("Insira o primeiro valor:")
		fmt.Scanf("%d", &valor1)
		fmt.Println("Insira o segundo valor:")
		fmt.Scanf("%d", &valor2)

		fmt.Println("Resultado: ", atividade.Divisao(valor1, valor2))

	case 5:

		fmt.Println(atividade2.Ponteiro())

	case 6:

		fmt.Println("Cadastro")
		fmt.Println("Digite o nome:")
		fmt.Scanf("%s", &nome)
		fmt.Println("Digite o email:")
		fmt.Scanf("%s", &email)

		fmt.Println(atividade2.CriarUsuario(nome, email))

	default:
		fmt.Println("Seção encerrada.")
	}

}
