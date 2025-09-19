//Este é um projeto simples de um gerenciador de tarefas em linha de comando (CLI) construído em Go. Ele permite adicionar novas tarefas e visualizar todas as tarefas salvas. O código permite que o usuário crie tarefas ou veja as tarefas criadas.

package main

import (
	"fmt"
)

func main() {
	var tarefas []string

	for {
		fmt.Println("\n--- Gerenciador de Tarefas ---")
		fmt.Println("1. Adicionar tarefa")
		fmt.Println("2. Ver todas as tarefas")
		fmt.Println("3. Sair")
		fmt.Print("Escolha uma opção: ")

		var escolha string
		fmt.Scanln(&escolha)

		switch escolha {
		case "1":
			fmt.Print("Digite a nova tarefa: ")
			var tarefa string
			fmt.Scanln(&tarefa)
			tarefas = append(tarefas, tarefa)
			fmt.Println("Tarefa adicionada com sucesso!")
		case "2":
			if len(tarefas) == 0 {
				fmt.Println("Nenhuma tarefa encontrada.")
				continue
			}
			fmt.Println("\n--- Suas Tarefas ---")
			for i, tarefa := range tarefas {
				fmt.Printf("%d. %s\n", i+1, tarefa)
			}
			fmt.Println("----------------------\n")
		case "3":
			fmt.Println("Até mais!")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
