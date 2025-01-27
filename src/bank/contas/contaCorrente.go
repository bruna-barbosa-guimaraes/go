package contas

import "bank/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	if valorSaque > 0 && valorSaque > c.saldo {
		return "saldo insuficiente. saldo atual: R$", c.saldo
	}
	c.saldo -= valorSaque
	return "Saque realizado com sucesso. saldo atual: R$", c.saldo
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso. saldo atual: R$", c.saldo
	} else {
		return "Valor do depósito menor ou igual a zero. saldo atual: R$", c.saldo
	}
}

func (contaOrigem *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) (string, float64) {
	if valorTransferencia > 0 && valorTransferencia < contaOrigem.saldo {
		contaOrigem.saldo -= valorTransferencia
		contaDestino.saldo += valorTransferencia
		return "Transferência realizada com sucesso. saldo atual: R$", contaOrigem.saldo
	} else {
		return "saldo insuficiente. saldo atual: R$", contaOrigem.saldo
	}

}

func (c *ContaCorrente) VerSaldo() float64 {
	return c.saldo
}
