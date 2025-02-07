package repository

import (
	"github.com/8soat-grupo35/fastfood-order-production/internal/entities"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetPaymentStatus(orderID uint32) (string, error)
	UpdatePaymentStatus(orderID uint32, status string) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(orm *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: orm,
	}
}

func (r *paymentRepository) GetPaymentStatus(orderID uint32) (string, error) {
	var payment entities.Payment
	if err := r.db.First(&payment, orderID).Error; err != nil {
		return "", err
	}
	return payment.Status, nil
}

func (r *paymentRepository) UpdatePaymentStatus(orderID uint32, status string) error {
	return r.db.Model(&entities.Payment{}).Where("id = ?", orderID).Update("status", status).Error
}