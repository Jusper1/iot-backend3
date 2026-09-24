package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type IOTInaprocHandler struct {
	Service *services.IOTInaprocService
}

func NewIOTInaprocHandler(
	service *services.IOTInaprocService,
) *IOTInaprocHandler {
	return &IOTInaprocHandler{
		Service: service,
	}
}

func (h *IOTInaprocHandler) FindAll(c *gin.Context) {
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

func (h *IOTInaprocHandler) FindByID(c *gin.Context) {
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
			"error":   "data IOT INAPROC tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *IOTInaprocHandler) FindByKode(c *gin.Context) {
	kode := c.Param("kode")

	data, err := h.Service.FindByKode(kode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "data IOT INAPROC tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *IOTInaprocHandler) Create(c *gin.Context) {

	var req models.IOTInaprocCreateRequest

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
		"message": "IOT INAPROC berhasil dibuat",
		"data":    data,
	})
}

func (h *IOTInaprocHandler) Update(c *gin.Context) {
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
		"message": "IOT INAPROC berhasil diperbarui",
		"data":    data,
	})
}

func (h *IOTInaprocHandler) Delete(c *gin.Context) {
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
		"message": "IOT INAPROC berhasil dihapus",
	})
}