package contas

import c "gobank/clientes"

type ContaPoupanca struct {
	Titular                              c.Cliente
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {

	if valorSaque > 0 && valorSaque <= c.saldo {
		c.saldo -= valorSaque
		return "Saque realizado!"
	} else {
		return "saldo insuficiente!"
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado! Seu novo saldo é:", c.saldo
	} else {
		return "Valor precisa ser maior que 0! O valor informado foi", valorDeposito
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
