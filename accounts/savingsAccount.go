package accounts

import c "gobank/holders"

type SavingAccount struct {
	Holder                                 c.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) Withdraw(withdrawValue float64) string {

	if withdrawValue > 0 && withdrawValue <= c.balance {
		c.balance -= withdrawValue
		return "Withdraw succeed!"
	} else {
		return "Insufficient funds!"
	}
}

func (c *SavingAccount) Deposit(depositValue float64) (string, float64) {

	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit succeed! New balance:", c.balance
	} else {
		return "Value must be greater than 0! The value informed was:", depositValue
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
