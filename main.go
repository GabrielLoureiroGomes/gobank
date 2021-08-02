package main

import (
	"fmt"
	cc "gobank/clientes"
	c "gobank/contas"
)

func main() {

	jessica := cc.CriaCliente("Jéssica", "222.222.222-22", "Psicóloga")
	gabriel := cc.CriaCliente("Gabriel", "111.111.111-11", "Desenvolvedor")

	fmt.Println(jessica)
	fmt.Println(gabriel)

	contaCorrenteGabriel := c.CriaConta("Gabriel", 111, 111222, 500)
	contaCorrenteJessica := c.CriaConta("Jessica", 222, 333444, 500)

	fmt.Println(contaCorrenteGabriel.ObterSaldo())
	fmt.Println(contaCorrenteJessica.ObterSaldo())

	contaCorrenteGabriel.Depositar(2000)
	contaCorrenteJessica.Sacar(250)

	fmt.Println(contaCorrenteGabriel.ObterSaldo())
	fmt.Println(contaCorrenteJessica.ObterSaldo())

	contaCorrenteGabriel.Sacar(500)
	contaCorrenteJessica.Depositar(750)

	fmt.Println(contaCorrenteGabriel.ObterSaldo())
	fmt.Println(contaCorrenteJessica.ObterSaldo())

	contaCorrenteGabriel.Transferir(500, &contaCorrenteJessica)

	fmt.Println(contaCorrenteGabriel.ObterSaldo())
	fmt.Println(contaCorrenteJessica.ObterSaldo())

	contaPoupancaGabriel := c.ContaPoupanca{}
	fmt.Println(contaPoupancaGabriel)

	contaPoupancaGabriel.Depositar(1000)
	fmt.Println(contaPoupancaGabriel.ObterSaldo())
	PagarBoleto(&contaPoupancaGabriel, 900)
	fmt.Println(contaPoupancaGabriel.ObterSaldo())

}

func PagarBoleto(conta verificaConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificaConta interface {
	Sacar(valor float64) string
}
