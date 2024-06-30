package main

import (
	"fmt"
	"study_golang/atividade"
)

func main() {

	var valor1, valor2 int
	var operacao = 1

	for operacao != 0 {

		fmt.Println("Menu")
		fmt.Println("1-soma")
		fmt.Println("2-Subtração")
		fmt.Println("3-Multiplicação")
		fmt.Println("4-Divisão")
		fmt.Println("0-Sair")
		fmt.Scanf("%d", &operacao)

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

		default:
			fmt.Println("Seção encerrada.")
		}

	}

}
