package main

import (
	"bank/contas"
	"fmt"
)

func main() {
	// contaBruna := ContaCorrente{titular: "Bruna", numeroAgencia: 589, numeroConta: 123456, saldo: 500}
	// contaGuilherme := ContaCorrente{"Guilherme", 222, 654321, 225.5}
	// contaVinicius := ContaCorrente{titular: "Vinicius", numeroConta: 123456}
	// fmt.Println(contaBruna)
	// fmt.Println(contaGuilherme)
	// fmt.Println(contaVinicius)

	// var contaCris *ContaCorrente
	// contaCris = new(ContaCorrente)
	// contaCris.titular = "Cris"
	// fmt.Println(contaCris)

	// contaSilvia := ContaCorrente{titular: "Silvia", numeroAgencia: 589, numeroConta: 234567, saldo: 500}
	// fmt.Println(contaSilvia.saldo)
	// fmt.Println(contaSilvia.Sacar(200))
	// fmt.Println(contaSilvia.saldo)
	// fmt.Println(contaSilvia.Depositar(-100))
	// fmt.Println(contaSilvia.saldo)

	// contaSilvia := contas.ContaCorrente{Titular: "Silvia", NumeroAgencia: 589, NumeroConta: 234567, Saldo: 500}
	// contaGuilherme := contas.ContaCorrente{Titular: "Guilherme", NumeroAgencia: 222, NumeroConta: 123456, Saldo: 225.5}
	// fmt.Println(contaSilvia)
	// fmt.Println(contaGuilherme)
	// fmt.Println(contaSilvia.Transferir(200, &contaGuilherme))
	// fmt.Println(contaSilvia)
	// fmt.Println(contaGuilherme)

	contaBruna := contas.ContaPoupanca{}
	contaBruna.Depositar(200)
	PagarBoleto(&contaBruna, 60)
	fmt.Println(contaBruna.VerSaldo())

	contaVinicius := contas.ContaCorrente{}
	contaVinicius.Depositar(450)
	PagarBoleto(&contaVinicius, 100)
	fmt.Println(contaVinicius.VerSaldo())

}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valorSaque float64) (string, float64)
}
