package types

import (
	"github.com/sev-2/raiden"
)

type Addresses struct {
	raiden.TypeBase
}

func (t *Addresses) Name() string {
	return "_addresses"
}

func (r *Addresses) Format() string {
	return "addresses[]"
}

func (r *Addresses) Enums() []string {
	return []string{}
}

func (r *Addresses) Comment() *string {
	return nil
}

