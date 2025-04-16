package rpc

import (
	"github.com/sev-2/raiden"
)

type SetDefaultPatientRoleParams struct {
}
type SetDefaultPatientRoleResult interface{}

type SetDefaultPatientRole struct {
	raiden.RpcBase
	Params   *SetDefaultPatientRoleParams `json:"-"`
	Return   SetDefaultPatientRoleResult `json:"-"`
}

func (r *SetDefaultPatientRole) GetName() string {
	return "set_default_patient_role"
}

func (r *SetDefaultPatientRole) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeTrigger
}

func (r *SetDefaultPatientRole) GetRawDefinition() string {
	return `BEGIN IF NEW.role_id IS NULL THEN NEW.role_id := 'e465f5f4-5b96-4359-8bc3-28e167c55407'; -- Role ID untuk patient END IF; RETURN NEW; END;`
}