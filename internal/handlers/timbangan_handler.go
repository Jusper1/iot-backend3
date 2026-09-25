package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type TimbanganHandler struct {
	Service *services.TimbanganService
}

func NewTimbanganHandler(
	service *services.TimbanganService,
) *TimbanganHandler {
	return &TimbanganHandler{
		Service: service,
	}
}

func (h *TimbanganHandler) FindAll(c *gin.Context) {
	data, err := h.Service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *TimbanganHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "ID tidak valid",
		})
		return
	}

	data, err := h.Service.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "data timbangan/RCW tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *TimbanganHandler) FindByKode(c *gin.Context) {
	kode := c.Param("kode")

	data, err := h.Service.FindByKode(kode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "data timbangan/RCW tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *TimbanganHandler) Create(c *gin.Context) {
	var req models.TimbanganCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "format JSON tidak valid",
		})
		return
	}

	data, err := h.Service.CreateFull(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Order timbangan/RCW berhasil dibuat",
		"data":    data,
	})
}

func (h *TimbanganHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "ID tidak valid",
		})
		return
	}

	var body map[string]interface{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "format JSON tidak valid",
		})
		return
	}

	if len(body) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "data update wajib diisi",
		})
		return
	}

	if err := h.Service.Update(uint(id), body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	data, err := h.Service.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Order timbangan/RCW berhasil diperbarui",
		"data":    data,
	})
}

func (h *TimbanganHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "ID tidak valid",
		})
		return
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Order timbangan/RCW berhasil dihapus",
	})
}