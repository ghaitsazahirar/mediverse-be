package types

import (
	"github.com/sev-2/raiden"
)

type EPrescriptions struct {
	raiden.TypeBase
}

func (t *EPrescriptions) Name() string {
	return "_e_prescriptions"
}

func (r *EPrescriptions) Format() string {
	return "e_prescriptions[]"
}

func (r *EPrescriptions) Enums() []string {
	return []string{}
}

func (r *EPrescriptions) Comment() *string {
	return nil
}

