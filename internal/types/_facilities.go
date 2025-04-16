package types

import (
	"github.com/sev-2/raiden"
)

type Facilities struct {
	raiden.TypeBase
}

func (t *Facilities) Name() string {
	return "_facilities"
}

func (r *Facilities) Format() string {
	return "facilities[]"
}

func (r *Facilities) Enums() []string {
	return []string{}
}

func (r *Facilities) Comment() *string {
	return nil
}

