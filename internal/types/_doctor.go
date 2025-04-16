package types

import (
	"github.com/sev-2/raiden"
)

type Doctor struct {
	raiden.TypeBase
}

func (t *Doctor) Name() string {
	return "_doctor"
}

func (r *Doctor) Format() string {
	return "doctor[]"
}

func (r *Doctor) Enums() []string {
	return []string{}
}

func (r *Doctor) Comment() *string {
	return nil
}

