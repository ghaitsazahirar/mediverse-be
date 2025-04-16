package types

import (
	"github.com/sev-2/raiden"
)

type Poli struct {
	raiden.TypeBase
}

func (t *Poli) Name() string {
	return "_poli"
}

func (r *Poli) Format() string {
	return "poli[]"
}

func (r *Poli) Enums() []string {
	return []string{}
}

func (r *Poli) Comment() *string {
	return nil
}

