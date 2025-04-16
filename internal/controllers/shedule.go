package controllers

import (
	"mediversebe/internal/models"

	"strings"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type ScheduleControllers struct {
	raiden.ControllerBase
	Http  string `path:"/schedules" type:"rest"`
	Model models.Schedules
}

func (c *ScheduleControllers) BeforeGet(ctx raiden.Context) error {
	authHeader := ctx.RequestContext().Request.Header.Peek("Authorization")
	apiKey := ctx.RequestContext().Request.Header.Peek("apikey")

	if len(authHeader) == 0 || !strings.HasPrefix(string(authHeader), "Bearer ") {
		return ctx.SendJson(map[string]string{
			"message": "Unauthorized",
			"hint":    "Missing or malformed Authorization header",
		})
	}

	if len(apiKey) == 0 {
		return ctx.SendJson(map[string]string{
			"message": "No API key found in request",
			"hint":    "No `apikey` request header or url param was found.",
		})
	}

	if string(apiKey) != "your-expected-api-key" {
		return ctx.SendJson(map[string]string{
			"message": "Invalid API key",
		})
	}

	return nil
}

func (c *ScheduleControllers) Get(ctx raiden.Context) error {
	query := db.NewQuery(ctx).From(models.Schedules{})
	params := ctx.RequestContext().QueryArgs()

	params.VisitAll(func(key, value []byte) {
		paramKey := string(key)
		paramVal := string(value)

		switch {
		case strings.HasPrefix(paramVal, "eq."):
			query = query.Eq(paramKey, strings.TrimPrefix(paramVal, "eq."))
		case strings.HasPrefix(paramVal, "neq."):
			query = query.Neq(paramKey, strings.TrimPrefix(paramVal, "neq."))
		}
	})

	return ctx.SendJson(query)
}
