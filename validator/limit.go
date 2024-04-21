package validator

import (
	"fmt"

	"github.com/nellystanford/article/chain_of_responsibility/contract"
)

// OperationalLimitsValidator verifies if account has enough limit for operation.
type OperationalLimitsValidator struct {
	Next Validator
}

func (olv *OperationalLimitsValidator) SetNext(next Validator) {
	olv.Next = next
}

func (olv *OperationalLimitsValidator) Validate(transaction *contract.Transaction) error {
	if !AccountHasLimit() {
		return fmt.Errorf("cannot complete transaction because account doesn't have enough limit")
	}

	if olv.Next != nil {
		return olv.Next.Validate(transaction)
	}

	return nil
}

func AccountHasLimit() bool {
	// Implementation of limit analysis.
	return true
}
