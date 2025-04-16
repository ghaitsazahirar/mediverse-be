package rpc

import (
	"github.com/sev-2/raiden"
)

type SetUserRoleBasedOnEmailParams struct {
}
type SetUserRoleBasedOnEmailResult interface{}

type SetUserRoleBasedOnEmail struct {
	raiden.RpcBase
	Params   *SetUserRoleBasedOnEmailParams `json:"-"`
	Return   SetUserRoleBasedOnEmailResult `json:"-"`
}

func (r *SetUserRoleBasedOnEmail) GetName() string {
	return "set_user_role_based_on_email"
}

func (r *SetUserRoleBasedOnEmail) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeTrigger
}

func (r *SetUserRoleBasedOnEmail) GetRawDefinition() string {
	return `BEGIN IF NEW.email = 'superadmin@example.com' THEN NEW.role_id := 'bd112dbe-bd53-4486-9fa0-28d4c96f4120'; -- Super Admin ELSIF NEW.email = 'admin@example.com' THEN NEW.role_id := '352bd099-9156-4d6e-8e7c-f085d0b89e59'; -- Admin ELSIF NEW.email LIKE '%@mediverse.com' THEN NEW.role_id := 'aaab46fb-445e-4d96-9376-a5b27127aaa2'; -- Doctor 'aaab46fb-445e-4d96-9376-a5b27127aaa2' ELSE NEW.role_id := 'e465f5f4-5b96-4359-8bc3-28e167c55407'; -- Default role untuk Patient END IF; RETURN NEW; END;`
}