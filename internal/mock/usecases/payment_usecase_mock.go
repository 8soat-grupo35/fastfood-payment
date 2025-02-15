package mock

import "github.com/stretchr/testify/mock"

// MockPaymentUsecase é uma implementação mock do PaymentUsecase
type MockPaymentUsecase struct {
    mock.Mock
}

func (m *MockPaymentUsecase) GetPaymentStatus(orderID uint32) (string, error) {
    args := m.Called(orderID)
    return args.String(0), args.Error(1)
}

func (m *MockPaymentUsecase) UpdatePaymentStatus(orderID uint32, status string) error {
    args := m.Called(orderID, status)


    if args.Get(0) == nil {
        return nil
    }

    return args.Get(0).(error)
}