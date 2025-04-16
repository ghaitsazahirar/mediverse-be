package controllers

import (
	"encoding/json"
	"mediversebe/internal/middleware"
	"mediversebe/internal/models"
	"strings"

	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/db"
)

type DoctorController struct {
	raiden.ControllerBase
	Http  string `path:"/doctors" type:"rest"`
	Model models.Doctor
}

// BeforeGet akan dipanggil sebelum method Get, pastikan middleware dijalankan sebelum data diambil
func (c *DoctorController) BeforeGet(ctx raiden.Context) error {
	// Menggunakan middleware Auth
	if err := middleware.AuthMiddleware()(ctx); err != nil {
		return err
	}

	// Mengecek API Key setelah autentikasi berhasil
	apiKey := ctx.RequestContext().Request.Header.Peek("apikey")
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

// Get untuk mengambil data dokter
func (c *DoctorController) Get(ctx raiden.Context) error {
	// Menggunakan query untuk mendapatkan data dokter
	query := db.NewQuery(ctx).From(models.Doctor{})
	params := ctx.RequestContext().QueryArgs()

	// Memproses parameter query
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

	// Mengambil data dokter dari database
	var doctors []models.Doctor
	if err := query.Get(&doctors); err != nil {
		return ctx.SendJson(map[string]string{
			"message": "Error fetching data",
			"error":   err.Error(),
		})
	}

	return ctx.SendJson(doctors)
}

// Post untuk menyimpan data dokter baru
func (c *DoctorController) Post(ctx raiden.Context) error {
	var doctor models.Doctor
	body := ctx.RequestContext().Request.Body()

	// Mengikat data JSON dari body request
	if err := json.Unmarshal(body, &doctor); err != nil {
		return ctx.SendJson(map[string]string{
			"message": "Error binding data",
			"error":   err.Error(),
		})
	}

	// Simpan data dokter ke dalam database
	if err := db.NewQuery(ctx).Insert("doctors", &doctor); err != nil {
		return ctx.SendJson(map[string]string{
			"message": "Error saving doctor",
			"error":   err.Error(),
		})
	}

	// Mengirimkan pesan sukses jika berhasil disimpan
	return ctx.SendJson(map[string]string{
		"message": "Doctor created successfully",
	})
}
