package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var ErrInvalidPayment = errors.New("data pembayaran tidak valid")

type PaymentService struct {
	Repo *repositories.PaymentRepository
}

func NewPaymentService(
	repo *repositories.PaymentRepository,
) *PaymentService {
	return &PaymentService{
		Repo: repo,
	}
}

func (s *PaymentService) Create(data *models.Payment) error {
	if data.OrderID == 0 {
		return ErrInvalidPayment
	}

	if data.JumlahUangMasuk < 0 {
		return ErrInvalidPayment
	}

	return s.Repo.Create(data)
}

func (s *PaymentService) FindByID(id uint) (*models.Payment, error) {
	if id == 0 {
		return nil, ErrInvalidPayment
	}

	return s.Repo.FindByID(id)
}

func (s *PaymentService) FindByOrderID(orderID uint) ([]models.Payment, error) {
	if orderID == 0 {
		return nil, ErrInvalidPayment
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *PaymentService) FindAll() ([]models.Payment, error) {
	return s.Repo.FindAll()
}

func (s *PaymentService) Update(data *models.Payment) error {
	if data.ID == 0 || data.OrderID == 0 {
		return ErrInvalidPayment
	}

	if data.JumlahUangMasuk < 0 {
		return ErrInvalidPayment
	}

	return s.Repo.Update(data)
}

func (s *PaymentService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidPayment
	}

	return s.Repo.Delete(id)
}