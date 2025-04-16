// /internal/middleware/auth.go
package middleware

import (
	"strings"

	"github.com/sev-2/raiden"
)

// Autentikasi Middleware
func AuthMiddleware() func(ctx raiden.Context) error {
	return func(ctx raiden.Context) error {
		// Mengecek apakah ada header Authorization
		authHeader := ctx.RequestContext().Request.Header.Peek("Authorization")
		if len(authHeader) == 0 || !strings.HasPrefix(string(authHeader), "Bearer ") {
			return ctx.SendJson(map[string]string{
				"message": "Unauthorized",
				"hint":    "Missing or malformed Authorization header",
			})
		}

		// Lanjutkan ke handler berikutnya jika autentikasi valid
		return nil
	}
}
