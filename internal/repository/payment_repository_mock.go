package repository

import (
	"github.com/8soat-grupo35/fastfood-payment/internal/entities"
	"github.com/stretchr/testify/mock"
)

// MockPaymentRepository é uma implementação mock do PaymentRepository
type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) GetPaymentStatus(orderID uint32) (string, error) {
	args := m.Called(orderID)
	return args.String(0), args.Error(1)
}

func (m *MockPaymentRepository) UpdatePaymentStatus(orderID uint32, status string) error {
	args := m.Called(orderID, status)
	return args.Error(0)
}

func (m *MockPaymentRepository) Create(payment entities.Payment) (*entities.Payment, error) {
	args := m.Called(payment)
	return args.Get(0).(*entities.Payment), args.Error(1)
}
