package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	boasVindas()

	for {
		exibeComandos()

		comando := leComando()

		listaDeSites := []string{"Insira os sites que deseja monitorar aqui"}

		switch comando {
		case 1:
			fmt.Println("Monitorando...")
			for j := 0; j < 5; j++ {
				monitora(listaDeSites)
				time.Sleep(5 * time.Second)
				fmt.Println("")

			}
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}
