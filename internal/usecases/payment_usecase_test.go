package usecase

import (
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"

    "github.com/8soat-grupo35/fastfood-order-production/internal/entities"
    "github.com/8soat-grupo35/fastfood-order-production/internal/repository"
)

type PaymentUsecaseSuite struct {
    suite.Suite
    mockRepo *repository.MockPaymentRepository
    usecase  PaymentUsecase
    payment  *entities.Payment
}

func (suite *PaymentUsecaseSuite) SetupSuite() {
    suite.mockRepo = new(repository.MockPaymentRepository)
    suite.usecase = &paymentUsecase{
        paymentRepo: suite.mockRepo,
    }

    suite.payment = &entities.Payment{
        ID:     1,
        OrderID: 123,
        Status: "PAID",
    }
}

func (suite *PaymentUsecaseSuite) AfterTest(_, _ string) {
    suite.mockRepo.AssertExpectations(suite.T())
}

// Teste de sucesso: GetPaymentStatus retorna o status corretamente
// func (suite *PaymentUsecaseSuite) TestGetPaymentStatus_Success() {
//     suite.mockRepo.On("GetPaymentStatus", suite.payment.ID).Return(suite.payment.Status, nil)

//     status, err := suite.usecase.GetPaymentStatus(suite.payment.ID)
//     assert.NoError(suite.T(), err)
//     assert.Equal(suite.T(), suite.payment.Status, status)
// }

// Teste de erro: GetPaymentStatus falha ao buscar o pagamento
func (suite *PaymentUsecaseSuite) TestGetPaymentStatus_Error() {
    suite.mockRepo.On("GetPaymentStatus", suite.payment.ID).Return("", errors.New("record not found"))

    status, err := suite.usecase.GetPaymentStatus(suite.payment.ID)
    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), "", status)
}

// Teste de sucesso: UpdatePaymentStatus atualiza o status corretamente
func (suite *PaymentUsecaseSuite) TestUpdatePaymentStatus_Success() {
    suite.mockRepo.On("UpdatePaymentStatus", suite.payment.ID, "COMPLETED").Return(nil)

    err := suite.usecase.UpdatePaymentStatus(suite.payment.ID, "COMPLETED")
    assert.NoError(suite.T(), err)
}

// Teste de erro: UpdatePaymentStatus falha ao atualizar o status
func (suite *PaymentUsecaseSuite) TestUpdatePaymentStatus_Error() {
    suite.mockRepo.On("UpdatePaymentStatus", suite.payment.ID, "FAILED").Return(errors.New("update error"))

    err := suite.usecase.UpdatePaymentStatus(suite.payment.ID, "FAILED")
    assert.Error(suite.T(), err)
}

// Função para rodar a test suite
func TestPaymentUsecaseSuite(t *testing.T) {
    suite.Run(t, new(PaymentUsecaseSuite))
}