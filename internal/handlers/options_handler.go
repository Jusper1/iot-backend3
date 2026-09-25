package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/rules"
)

func GetDropdownOptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"jenis_kertas":    rules.JenisKertasOptions,
			"tipe_timbangan":  rules.TipeTimbanganOptions,
			"kode_bayar":      rules.KodeBayarOptions,
			"status_pesanan":  rules.StatusPesananOptions,
			"status_odoo":     rules.StatusOdooOptions,
			"jenis_file":      rules.JenisFileOptions,
		},
	})
}