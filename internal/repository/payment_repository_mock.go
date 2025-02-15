package repository

import "github.com/stretchr/testify/mock"

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