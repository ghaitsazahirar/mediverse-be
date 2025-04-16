package types

import (
	"github.com/sev-2/raiden"
)

type Patients struct {
	raiden.TypeBase
}

func (t *Patients) Name() string {
	return "_patients"
}

func (r *Patients) Format() string {
	return "patients[]"
}

func (r *Patients) Enums() []string {
	return []string{}
}

func (r *Patients) Comment() *string {
	return nil
}

