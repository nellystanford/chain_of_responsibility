package validator

import (
	"fmt"

	"github.com/nellystanford/article/chain_of_responsibility/contract"
)

// AntiFraudValidator verifies if transaction is fraudulent.
type AntiFraudValidator struct {
	Next Validator
}

func (afv *AntiFraudValidator) SetNext(next Validator) {
	afv.Next = next
}

func (afv *AntiFraudValidator) Validate(transaction *contract.Transaction) error {
	if isFraudulent() {
		return fmt.Errorf("cannot complete transaction because it was reproved by antifraud validation")
	}

	if afv.Next != nil {
		return afv.Next.Validate(transaction)
	}

	return nil
}

func isFraudulent() bool {
	// Implementation of antifraud analysis.
	return false
}
