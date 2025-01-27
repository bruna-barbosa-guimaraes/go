package contas

import "bank/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) (string, float64) {
	if valorSaque > 0 && valorSaque > c.saldo {
		return "saldo insuficiente. saldo atual: R$", c.saldo
	}
	c.saldo -= valorSaque
	return "Saque realizado com sucesso. saldo atual: R$", c.saldo
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso. saldo atual: R$", c.saldo
	} else {
		return "Valor do depósito menor ou igual a zero. saldo atual: R$", c.saldo
	}
}

func (c *ContaPoupanca) VerSaldo() float64 {
	return c.saldo
}
