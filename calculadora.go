//Uma calculadora simples feita em Go que interpreta expressões digitadas diretamente no terminal. Este projeto demonstra o uso de lógica de string para realizar operações matemáticas.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fatorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	for {
		fmt.Println("\n--- Calculadora de Linha de Comando ---")
		fmt.Println("Operadores suportados:| + | - | * | / | ! |")
		fmt.Println("Digite a sua expressão ou 'sair' para fechar.")
		fmt.Print("-> ")

		var expressao string
		fmt.Scanln(&expressao)

		if strings.ToLower(expressao) == "sair" {
			fmt.Println("Até logo!")
			os.Exit(0)
		}

		expressao = strings.ReplaceAll(expressao, " ", "")

		var resultado float64
		var erro bool
		var operacao string

		if strings.HasSuffix(expressao, "!") {
			operacao = "!"
			numStr := strings.TrimSuffix(expressao, "!")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				erro = true
			} else {
				resFatorial := fatorial(num)
				if resFatorial == 0 && num < 0 {
					fmt.Println("Erro: Fatorial de número negativo não é suportado.")
					continue
				}
				fmt.Printf("Resultado: %d\n", resFatorial)
			}
			continue
		}

		operadores := "+-*/^"
		partes := []string{}
		for _, op := range operadores {
			if strings.Contains(expressao, string(op)) {
				operacao = string(op)
				partes = strings.Split(expressao, operacao)
				break
			}
		}

		if len(partes) != 2 {
			fmt.Println("Erro: Expressão inválida. Verifique o formato.")
			erro = true
		} else {
			num1, err1 := strconv.ParseFloat(partes[0], 64)
			num2, err2 := strconv.ParseFloat(partes[1], 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Erro: Um ou mais valores são inválidos.")
				erro = true
			} else {
				switch operacao {
				case "+":
					resultado = num1 + num2
				case "-":
					resultado = num1 - num2
				case "*":
					resultado = num1 * num2
				case "/":
					if num2 == 0 {
						fmt.Println("Erro: Divisão por zero.")
						erro = true
					} else {
						resultado = num1 / num2
					}
				case "^":
					fmt.Println("Erro: Potenciação não implementada com estes recursos.")
					erro = true
				default:
					erro = true
				}
			}
		}

		if !erro {
			fmt.Printf("Resultado: %.2f\n", resultado)
		}
	}
}
