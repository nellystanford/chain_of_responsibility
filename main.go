package main

import (
	"fmt"

	"github.com/nellystanford/article/chain_of_responsibility/contract"
	"github.com/nellystanford/article/chain_of_responsibility/validator"
)

func main() {
	// Setting the chain of validators.
	accountStatus := &validator.AccountStatusValidator{}
	accountBalance := &validator.AccountBalanceValidator{}
	opLimits := &validator.OperationalLimitsValidator{}
	antifraud := &validator.AntiFraudValidator{}

	accountStatus.SetNext(accountBalance)
	accountBalance.SetNext(opLimits)
	opLimits.SetNext(antifraud)

	// Simulating a transaction.
	transaction := &contract.Transaction{
		Amount:        1000.0,
		AccountNumber: 2109,
	}

	// Validating the transaction starting on the first validator of the chain.
	err := accountStatus.Validate(transaction)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transaction validated successfully!")
}
