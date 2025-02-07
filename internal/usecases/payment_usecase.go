package usecase

import (
	"github.com/8soat-grupo35/fastfood-order-production/internal/repository"
	"gorm.io/gorm"
)

type PaymentUsecase interface {
	GetPaymentStatus(orderID uint32) (string, error)
	UpdatePaymentStatus(orderID uint32, status string) error
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentUsecase(db *gorm.DB) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: repository.NewPaymentRepository(db),
	}
}

func (u *paymentUsecase) GetPaymentStatus(orderID uint32) (string, error) {
	return u.paymentRepo.GetPaymentStatus(orderID)
}

func (u *paymentUsecase) UpdatePaymentStatus(orderID uint32, status string) error {
	return u.paymentRepo.UpdatePaymentStatus(orderID, status)
}