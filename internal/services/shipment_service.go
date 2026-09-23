package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var ErrInvalidShipment = errors.New("data pengiriman tidak valid")

type ShipmentService struct {
	Repo *repositories.ShipmentRepository
}

func NewShipmentService(
	repo *repositories.ShipmentRepository,
) *ShipmentService {
	return &ShipmentService{
		Repo: repo,
	}
}

func (s *ShipmentService) Create(data *models.Shipment) error {
	if data.OrderID == 0 {
		return ErrInvalidShipment
	}

	if data.Berat < 0 || data.Ongkir < 0 {
		return ErrInvalidShipment
	}

	if data.Resi != nil {
	resi := strings.TrimSpace(*data.Resi)
	data.Resi = &resi
	}

	return s.Repo.Create(data)
}

func (s *ShipmentService) FindByID(id uint) (*models.Shipment, error) {
	if id == 0 {
		return nil, ErrInvalidShipment
	}

	return s.Repo.FindByID(id)
}

func (s *ShipmentService) FindByOrderID(orderID uint) ([]models.Shipment, error) {
	if orderID == 0 {
		return nil, ErrInvalidShipment
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *ShipmentService) FindAll() ([]models.Shipment, error) {
	return s.Repo.FindAll()
}

func (s *ShipmentService) Update(data *models.Shipment) error {
	if data.ID == 0 || data.OrderID == 0 {
		return ErrInvalidShipment
	}

	if data.Berat < 0 || data.Ongkir < 0 {
		return ErrInvalidShipment
	}

	if data.Resi != nil {
	resi := strings.TrimSpace(*data.Resi)
	data.Resi = &resi
	}

	return s.Repo.Update(data)
}

func (s *ShipmentService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidShipment
	}

	return s.Repo.Delete(id)
}