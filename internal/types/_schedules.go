package types

import (
	"github.com/sev-2/raiden"
)

type Schedules struct {
	raiden.TypeBase
}

func (t *Schedules) Name() string {
	return "_schedules"
}

func (r *Schedules) Format() string {
	return "schedules[]"
}

func (r *Schedules) Enums() []string {
	return []string{}
}

func (r *Schedules) Comment() *string {
	return nil
}

