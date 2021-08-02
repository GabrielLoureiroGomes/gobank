package accounts

import c "gobank/holders"

type CheckingAccount struct {
	Holder                      c.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func CreateAccount(name string, AgencyNumber int, AccountNumber int, balance float64) CheckingAccount {
	Holder := c.CreateHolder("Gabriel", "111.111.111-11", "Developer")
	newAccount := CheckingAccount{}
	newAccount.Holder = Holder
	newAccount.AgencyNumber = AgencyNumber
	newAccount.AccountNumber = AccountNumber
	newAccount.balance = balance
	return newAccount
}

func (c *CheckingAccount) Withdraw(withdrawValue float64) string {

	if withdrawValue > 0 && withdrawValue <= c.balance {
		c.balance -= withdrawValue
		return "Withdraw succeed!"
	} else {
		return "Insufficient funds!"
	}
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {

	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit succeed! New balance:", c.balance
	} else {
		return "Value must be greater than 0! The value informed was:", depositValue
	}
}

func (c *CheckingAccount) Transfer(tranferValue float64, destinationAccount *CheckingAccount) string {

	if tranferValue > 0 && tranferValue <= c.balance {
		c.balance -= tranferValue
		destinationAccount.Deposit(tranferValue)

		return "Tranfer succeed!"
	} else {
		return "Value must be greater than 0!"
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
