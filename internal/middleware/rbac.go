package middleware

import (
	"mediversebe/internal/models"
	"net/http"

	"github.com/sev-2/raiden"
)

// RoleBasedAccessControl memeriksa role_id dari user berdasarkan role yang dibutuhkan
func RoleBasedAccessControl(requiredRoles ...string) func(raiden.Context) error {
	return func(ctx raiden.Context) error {
		// Ambil user yang sedang login (misalnya dari session atau token)
		userInterface := ctx.Get("user")
		if userInterface == nil {
			return ctx.SendJson(map[string]any{
				"status":  http.StatusUnauthorized,
				"message": "User not authorized",
			})
		}

		// Pastikan type assertion sukses
		user, ok := userInterface.(*models.Users)
		if !ok {
			return ctx.SendJson(map[string]any{
				"status":  http.StatusUnauthorized,
				"message": "User data is not valid",
			})
		}

		// Cek apakah role user ada dalam list requiredRoles
		roleAllowed := false
		for _, role := range requiredRoles {
			// Bandingkan role berdasarkan role_id yang ada pada user
			if user.RoleId.String() == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			return ctx.SendJson(map[string]any{
				"status":  http.StatusForbidden,
				"message": "Insufficient permissions",
			})
		}

		// Lanjutkan ke handler berikutnya jika role user valid
		return nil
	}
}
