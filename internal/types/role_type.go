package types

import (
	"github.com/sev-2/raiden"
)

type RoleType struct {
	raiden.TypeBase
}

func (t *RoleType) Name() string {
	return "role_type"
}

func (r *RoleType) Format() string {
	return "role_type"
}

func (r *RoleType) Enums() []string {
	return []string{"Super Admin","Admin","Doctors","Patients"}
}

func (r *RoleType) Comment() *string {
	return nil
}

