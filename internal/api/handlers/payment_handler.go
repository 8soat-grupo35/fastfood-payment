package handlers

import (
	"net/http"
	"strconv"

	usecase "github.com/8soat-grupo35/fastfood-order-production/internal/usecases"
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
// @Param		 orderID             path int         true "ID do pedido"
// @Router       /v1/orders/{orderID}/payment/status [get]
// @success 200 {object} presenters.OrderPaymentStatusPresenter
// @Failure 500 {object} error
func (h *PaymentHandler) GetPaymentStatus(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	status, err := h.paymentUsecase.GetPaymentStatus(uint32(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, status)
}

// UpdateOrderPaymentStatus godoc
// @Summary      Update Order Payment Status
// @Description  Update Order Payment Status
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param		 orderID             path int         true "ID do pedido"
// @Param        UpdateBody	body dto.OrderPaymentStatusDto true "UpdateBody"
// @Router       /v1/orders/{orderID}/payment/status [put]
// @success 200 {object} string
// @Failure 500 {object} error
func (h *PaymentHandler) UpdatePaymentStatus(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))
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