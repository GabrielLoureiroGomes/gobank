package contas

import c "gobank/clientes"

type ContaCorrente struct {
	Titular                    c.Cliente
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func CriaConta(nome string, NumeroAgencia int, NumeroConta int, saldo float64) ContaCorrente {
	titular := c.CriaCliente("Gabriel", "111.111.111-11", "Desenvolvedor")
	novaConta := ContaCorrente{}
	novaConta.Titular = titular
	novaConta.NumeroAgencia = NumeroAgencia
	novaConta.NumeroConta = NumeroConta
	novaConta.saldo = saldo
	return novaConta
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {

	if valorSaque > 0 && valorSaque <= c.saldo {
		c.saldo -= valorSaque
		return "Saque realizado!"
	} else {
		return "saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado! Seu novo saldo é:", c.saldo
	} else {
		return "Valor precisa ser maior que 0! O valor informado foi", valorDeposito
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) string {

	if valorTransferencia > 0 && valorTransferencia <= c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)

		return "Tranferência realizada!"
	} else {
		return "Valor precisa ser maior que 0!"
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
