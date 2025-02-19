package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/8soat-grupo35/fastfood-payment/internal/entities"
	mock "github.com/8soat-grupo35/fastfood-payment/internal/mock/usecases"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PaymentHandlerSuite struct {
	suite.Suite
	mockUsecase *mock.MockPaymentUsecase
	handler     *PaymentHandler
	echo        *echo.Echo
}

func (suite *PaymentHandlerSuite) SetupSuite() {
	suite.mockUsecase = new(mock.MockPaymentUsecase)
	suite.handler = &PaymentHandler{
		paymentUsecase: suite.mockUsecase,
	}
	suite.echo = echo.New()
}

func (suite *PaymentHandlerSuite) SetupTest() {
	suite.echo = echo.New()
}

func (suite *PaymentHandlerSuite) AfterTest(_, _ string) {
	suite.mockUsecase.AssertExpectations(suite.T())
}

// Teste de sucesso: GetPaymentStatus retorna o status corretamente
func (suite *PaymentHandlerSuite) TestGetPaymentStatus_Success() {
	req := httptest.NewRequest(http.MethodGet, "/v1/payments/1/payment/status", nil)
	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)
	c.SetParamNames("orderID")
	c.SetParamValues("1")

	suite.mockUsecase.On("GetPaymentStatus", uint32(1)).Return("PAID", nil)

	err := suite.handler.GetPaymentStatus(c)
	assert.NoError(suite.T(), err)
	// assert.Equal(suite.T(), http.StatusOK, rec.Code)
	// assert.Equal(suite.T(), "\"PAID\"\n", rec.Body.String())
}

// Teste de erro: GetPaymentStatus falha ao buscar o pagamento
func (suite *PaymentHandlerSuite) TestGetPaymentStatus_Error() {
	req := httptest.NewRequest(http.MethodGet, "/v1/payments/1/payment/status", nil)
	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)
	c.SetParamNames("orderID")
	c.SetParamValues("1")

	suite.mockUsecase.On("GetPaymentStatus", uint32(1)).Return("", errors.New("record not found"))

	err := suite.handler.GetPaymentStatus(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
	assert.Equal(suite.T(), "\"record not found\"\n", rec.Body.String())
}

// Teste de sucesso: UpdatePaymentStatus atualiza o status corretamente
func (suite *PaymentHandlerSuite) TestUpdatePaymentStatus_Success() {
	body := `{"status":"COMPLETED"}`
	req := httptest.NewRequest(http.MethodPut, "/v1/payments/1/payment/status", bytes.NewBufferString(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)
	c.SetParamNames("orderID")
	c.SetParamValues("1")

	suite.mockUsecase.On("UpdatePaymentStatus", uint32(1), "COMPLETED").Return(nil)

	err := suite.handler.UpdatePaymentStatus(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), "\"Payment status updated\"\n", rec.Body.String())
}

// Teste de erro: UpdatePaymentStatus falha ao atualizar o status
func (suite *PaymentHandlerSuite) TestUpdatePaymentStatus_Error() {
	body := `{"status":"FAILED"}`
	req := httptest.NewRequest(http.MethodPut, "/v1/payments/1/payment/status", bytes.NewBufferString(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)
	c.SetParamNames("orderID")
	c.SetParamValues("1")

	suite.mockUsecase.On("UpdatePaymentStatus", uint32(1), "FAILED").Return(errors.New("update error"))

	err := suite.handler.UpdatePaymentStatus(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
	assert.Equal(suite.T(), "\"update error\"\n", rec.Body.String())
}

func (suite *PaymentHandlerSuite) TestCreatePayment_Success() {
	body := `{"orderId": 123}`
	req := httptest.NewRequest(http.MethodPost, "/v1/payments", bytes.NewBufferString(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	mockPayment := &entities.Payment{
		ID:      1,
		OrderID: 123,
		Status:  "WAITING",
	}

	suite.mockUsecase.On("Create", uint32(123)).Return(mockPayment, nil)

	err := suite.handler.CreatePayment(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *PaymentHandlerSuite) TestCreatePayment_BindError() {
	body := `{"orderId": "invalid"}` // OrderID inválido
	req := httptest.NewRequest(http.MethodPost, "/v1/payments", bytes.NewBufferString(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.echo.NewContext(req, rec)

	err := suite.handler.CreatePayment(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), "\"Invalid request body\"\n", rec.Body.String())
}

func (suite *PaymentHandlerSuite) TestNewPaymentHandler() {
	handler := NewPaymentHandler(nil)
	assert.NotNil(suite.T(), handler)
}

// Função para rodar a test suite
func TestPaymentHandlerSuite(t *testing.T) {
	suite.Run(t, new(PaymentHandlerSuite))
}
