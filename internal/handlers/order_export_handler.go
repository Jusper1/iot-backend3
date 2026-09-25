package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/models"
	"iot-backend/internal/rules"
	"iot-backend/internal/services"
)

type OrderExportHandler struct {
	Service *services.OrderExportService
}

func NewOrderExportHandler(
	service *services.OrderExportService,
) *OrderExportHandler {
	return &OrderExportHandler{
		Service: service,
	}
}

func (h *OrderExportHandler) Export(c *gin.Context) {
	filter := models.OrderExportFilter{
		KodeOrder:              strings.TrimSpace(c.Query("kode_order")),
		NamaInstansi:           strings.TrimSpace(c.Query("nama_instansi")),
		Status:                 strings.TrimSpace(c.Query("status")),
		Provinsi:               strings.TrimSpace(c.Query("provinsi")),
		TanggalPODari:          strings.TrimSpace(c.Query("tanggal_po_dari")),
		TanggalPOSampai:        strings.TrimSpace(c.Query("tanggal_po_sampai")),
		TanggalUangMasukDari:   strings.TrimSpace(c.Query("tanggal_uang_masuk_dari")),
		TanggalUangMasukSampai: strings.TrimSpace(c.Query("tanggal_uang_masuk_sampai")),
	}

	if rawKategori := c.Query("kategori_order"); rawKategori != "" {
		for _, k := range strings.Split(rawKategori, ",") {
			k = strings.TrimSpace(k)
			if k == "" {
				continue
			}
			if !rules.IsKategoriValid(k) {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   fmt.Sprintf("kategori_order tidak dikenal: %s", k),
				})
				return
			}
			filter.KategoriOrder = append(filter.KategoriOrder, k)
		}
	}

	for _, v := range []string{
		filter.TanggalPODari, filter.TanggalPOSampai,
		filter.TanggalUangMasukDari, filter.TanggalUangMasukSampai,
	} {
		if v == "" {
			continue
		}
		if _, err := time.Parse("2006-01-02", v); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   fmt.Sprintf("format tanggal tidak valid: %q, gunakan YYYY-MM-DD", v),
			})
			return
		}
	}

	fileBytes, err := h.Service.GenerateExcel(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	filename := fmt.Sprintf("data-pesanan-%s.xlsx", time.Now().Format("20060102-150405"))

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	c.Data(
		http.StatusOK,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		fileBytes,
	)
}