package entities

const (
	PAYMENT_STATUS_WAITING  = "WAITING"
	PAYMENT_STATUS_RECUSED  = "RECUSED"
	PAYMENT_STATUS_APPROVED = "APPROVED"
)

type Payment struct {
	ID      uint32 `gorm:"primaryKey"`
	OrderID uint32 `gorm:"not null"`
	Status  string `gorm:"not null"`
}

func NewOrderPayment(orderID uint32, status string) *Payment {
	orderPayment := &Payment{
		OrderID: orderID,
		Status:  status,
	}

	return orderPayment
}
