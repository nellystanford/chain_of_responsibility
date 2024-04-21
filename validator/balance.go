package validator

import (
	"fmt"

	"github.com/nellystanford/article/chain_of_responsibility/contract"
)

// AccountBalanceValidator verifies if account balance is enough for the transaction.
type AccountBalanceValidator struct {
	Next Validator
}

func (abv *AccountBalanceValidator) SetNext(next Validator) {
	abv.Next = next
}

func (abv *AccountBalanceValidator) Validate(transaction *contract.Transaction) error {
	if !AccountHasEnoughBalance() {
		return fmt.Errorf("cannot complete transaction because account doesn't have enough balance")
	}

	if abv.Next != nil {
		return abv.Next.Validate(transaction)
	}

	return nil
}

func AccountHasEnoughBalance() bool {
	// Implementation of balance analysis.
	return true
}
