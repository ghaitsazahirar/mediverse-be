package types

import (
	"github.com/sev-2/raiden"
)

type Services struct {
	raiden.TypeBase
}

func (t *Services) Name() string {
	return "_services"
}

func (r *Services) Format() string {
	return "services[]"
}

func (r *Services) Enums() []string {
	return []string{}
}

func (r *Services) Comment() *string {
	return nil
}

