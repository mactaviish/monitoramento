package main

import (
	"fmt"
)

func main() {
	fmt.Println("")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
	fmt.Print("Digite a operação desejada: ")

	var command int
	fmt.Scan(&command)

	switch command {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 0:
		exit()
	}
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento...")
}

func showLogs() {
	fmt.Println("Exibindo logs...")
}

func exit() {
	fmt.Println("Saindo...")
}
