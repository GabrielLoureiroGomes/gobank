package main

import (
	"fmt"
	c "gobank/accounts"
	cc "gobank/holders"
)

func main() {

	jessica := cc.CreateHolder("Jéssica", "222.222.222-22", "Psychologist")
	gabriel := cc.CreateHolder("Gabriel", "111.111.111-11", "Developer")

	fmt.Println(jessica)
	fmt.Println(gabriel)

	checkingAccountGabriel := c.CreateAccount("Gabriel", 111, 111222, 500)
	checkingAccountJessica := c.CreateAccount("Jessica", 222, 333444, 500)

	fmt.Println(checkingAccountGabriel.GetBalance())
	fmt.Println(checkingAccountJessica.GetBalance())

	checkingAccountGabriel.Deposit(2000)
	checkingAccountJessica.Withdraw(250)

	fmt.Println(checkingAccountGabriel.GetBalance())
	fmt.Println(checkingAccountJessica.GetBalance())

	checkingAccountGabriel.Withdraw(500)
	checkingAccountJessica.Deposit(750)

	fmt.Println(checkingAccountGabriel.GetBalance())
	fmt.Println(checkingAccountJessica.GetBalance())

	checkingAccountGabriel.Transfer(500, &checkingAccountJessica)

	fmt.Println(checkingAccountGabriel.GetBalance())
	fmt.Println(checkingAccountJessica.GetBalance())

	savingAccountGabriel := c.SavingAccount{}
	fmt.Println(savingAccountGabriel)

	savingAccountGabriel.Deposit(1000)
	fmt.Println(savingAccountGabriel.GetBalance())
	PayBillet(&savingAccountGabriel, 900)
	fmt.Println(savingAccountGabriel.GetBalance())

}

func PayBillet(account verifyAccount, billetValue float64) {
	account.Withdraw(billetValue)
}

type verifyAccount interface {
	Withdraw(withdrawValue float64) string
}
