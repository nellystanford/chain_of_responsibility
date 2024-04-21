package validator

import "github.com/nellystanford/article/chain_of_responsibility/contract"

// Validator defines interface for the chain validators.
type Validator interface {
	Validate(*contract.Transaction) error
	SetNext(Validator)
}
