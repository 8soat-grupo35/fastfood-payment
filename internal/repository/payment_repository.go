package repository

import (
	"log"

	"github.com/8soat-grupo35/fastfood-payment/internal/entities"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetPaymentStatus(orderID uint32) (string, error)
	UpdatePaymentStatus(orderID uint32, status string) error
	Create(payment entities.Payment) (*entities.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(orm *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: orm,
	}
}

func (repository *paymentRepository) GetPaymentStatus(orderID uint32) (string, error) {
	var payment entities.Payment
	if err := repository.db.First(&payment, orderID).Error; err != nil {
		return "", err
	}
	return payment.Status, nil
}

func (repository *paymentRepository) UpdatePaymentStatus(orderID uint32, status string) error {
	return repository.db.Model(&entities.Payment{}).Where("id = ?", orderID).Update("status", status).Error
}

func (repository *paymentRepository) Create(payment entities.Payment) (*entities.Payment, error) {
	result := repository.db.Create(&payment)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &payment, nil
}
