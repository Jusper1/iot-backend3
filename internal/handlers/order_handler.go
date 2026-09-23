package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"iot-backend/internal/models"
	"iot-backend/internal/services"
)

type OrderHandler struct {
	Service *services.OrderService
}

func NewOrderHandler(s *services.OrderService) *OrderHandler {
	return &OrderHandler{
		Service: s,
	}
}

func (h *OrderHandler) FindAll(c *gin.Context) {
	data, err := h.Service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "gagal mengambil data order",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *OrderHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID order tidak valid",
		})
		return
	}

	data, err := h.Service.FindByID(uint(id))
	if err != nil {
		status := http.StatusInternalServerError

		if errors.Is(err, services.ErrOrderNotFound) {
			status = http.StatusNotFound
		}

		c.JSON(status, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
func (h *OrderHandler) FindByKode(c *gin.Context) {
	kode := c.Param("kode")

	data, err := h.Service.FindByKode(kode)
	if err != nil {
		status := http.StatusInternalServerError

		if errors.Is(err, services.ErrOrderNotFound) {
			status = http.StatusNotFound
		}

		c.JSON(status, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *OrderHandler) Create(c *gin.Context) {
	var data models.Order

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "format request tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := h.Service.Create(&data); err != nil {
		status := http.StatusInternalServerError

		if errors.Is(err, services.ErrKodeOrderRequired) {
			status = http.StatusBadRequest
		}

		if errors.Is(err, services.ErrKodeOrderExists) {
			status = http.StatusConflict
		}

		c.JSON(status, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "order berhasil dibuat",
		"data":    data,
	})
}

func (h *OrderHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID order tidak valid",
		})
		return
	}

	var data models.Order

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "format request tidak valid",
			"error":   err.Error(),
		})
		return
	}

	data.ID = uint(id)

	if err := h.Service.Update(&data); err != nil {
		status := http.StatusInternalServerError

		switch {
		case errors.Is(err, services.ErrOrderNotFound):
			status = http.StatusNotFound
		case errors.Is(err, services.ErrKodeOrderRequired):
			status = http.StatusBadRequest
		case errors.Is(err, services.ErrKodeOrderExists):
			status = http.StatusConflict
		}

		c.JSON(status, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "order berhasil diperbarui",
		"data":    data,
	})
}

func (h *OrderHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID order tidak valid",
		})
		return
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		status := http.StatusInternalServerError

		if errors.Is(err, services.ErrOrderNotFound) {
			status = http.StatusNotFound
		}

		c.JSON(status, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "order berhasil dihapus",
	})
}