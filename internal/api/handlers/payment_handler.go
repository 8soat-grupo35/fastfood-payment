package handlers

import (
	"net/http"
	"strconv"

	usecase "github.com/8soat-grupo35/fastfood-payment/internal/usecases"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	paymentUsecase usecase.PaymentUsecase
}

func NewPaymentHandler(db *gorm.DB) *PaymentHandler {
	return &PaymentHandler{
		paymentUsecase: usecase.NewPaymentUsecase(db),
	}
}

// GetOrderPaymentStatus godoc
// @Summary      Get Order Payment Status
// @Description  Get Order Payment Status
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        orderID  path  int  true  "ID do pedido"
// @Router       /v1/payments/{orderID}/payment/status [get]
// @Success      200  {object}  string
// @Failure      500  {object}  error
func (h *PaymentHandler) GetPaymentStatus(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	status, err := h.paymentUsecase.GetPaymentStatus(uint32(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, status)
}

// UpdatePaymentStatus godoc
// @Summary      Update Order Payment Status
// @Description  Atualiza o status de pagamento de um pedido
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        orderID  path  int  true  "ID do pedido"
// @Param        status   body  string  true  "Status do pagamento"
// @Router       /v1/payments/{orderID}/payment/status [put]
// @Success      200  {object}  string
// @Failure      500  {object}  error
func (h *PaymentHandler) UpdatePaymentStatus(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	var request struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	err = h.paymentUsecase.UpdatePaymentStatus(uint32(orderID), request.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Payment status updated")
}

// CreatePayment godoc
// @Summary      Create Payment
// @Description  Create payment form order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        orderId   body  string  true  "Id do pedido"
// @Router       /v1/payments [post]
// @Success      200  {object}  int
// @Failure      500  {object}  error
func (h *PaymentHandler) CreatePayment(c echo.Context) error {
	var request struct {
		OrderID uint32 `json:"orderId"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	payment, err := h.paymentUsecase.Create(request.OrderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, payment)
}
