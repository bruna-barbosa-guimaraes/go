package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	if valorSaque > 0 && valorSaque > c.Saldo {
		return "Saldo insuficiente. Saldo atual: R$", c.Saldo
	}
	c.Saldo -= valorSaque
	return "Saque realizado com sucesso. Saldo atual: R$", c.Saldo
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "Depósito realizado com sucesso. Saldo atual: R$", c.Saldo
	} else {
		return "Valor do depósito menor ou igual a zero. Saldo atual: R$", c.Saldo
	}
}

func (contaOrigem *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) (string, float64) {
	if valorTransferencia > 0 && valorTransferencia < contaOrigem.Saldo {
		contaOrigem.Saldo -= valorTransferencia
		contaDestino.Saldo += valorTransferencia
		return "Transferência realizada com sucesso. Saldo atual: R$", contaOrigem.Saldo
	} else {
		return "Saldo insuficiente. Saldo atual: R$", contaOrigem.Saldo
	}

}
