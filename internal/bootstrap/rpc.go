// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden/pkg/resource"
	"mediversebe/internal/rpc"
)

func RegisterRpc() {
	resource.RegisterRpc(
		&rpc.SetDefaultPatientRole{},
		&rpc.SetUserRoleBasedOnEmail{},
	)
}
