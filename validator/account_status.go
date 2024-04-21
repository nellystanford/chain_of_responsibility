package validator

import (
	"fmt"

	"github.com/nellystanford/article/chain_of_responsibility/contract"
)

// AccountStatusValidator verifies the account status.
type AccountStatusValidator struct {
	Next Validator
}

func (asv *AccountStatusValidator) SetNext(next Validator) {
	asv.Next = next
}

func (asv *AccountStatusValidator) Validate(transaction *contract.Transaction) error {
	if !accountIsActive() {
		return fmt.Errorf("cannot complete transaction because account status is invalid")
	}

	if asv.Next != nil {
		return asv.Next.Validate(transaction)
	}

	return nil
}

func accountIsActive() bool {
	// Validation of account status.
	return true
}
