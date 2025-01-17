package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibirIntroducao()

	for {
		exibirMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}

func exibirIntroducao() {
	nome := "Bruna"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibirMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando foi:", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site := "https://www.alura.com.br"
	site := "https://httpbin.org/status/404"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
