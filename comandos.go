package main

import "fmt"

func exibeComandos() {

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido é", comando)
	fmt.Println("")
	return comando
}
