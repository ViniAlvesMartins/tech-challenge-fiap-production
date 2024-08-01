package entity

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"time"
)

type (
	OrderStatusUpdatedMessage struct {
		OrderId int              `json:"order_id"`
		Status  enum.OrderStatus `json:"status"`
	}

	PaymentStatusUpdatedMessage struct {
		OrderId int                `json:"order_id"`
		Status  enum.PaymentStatus `json:"status"`
	}

	Production struct {
		ID        string                `json:"id"`
		OrderId   int                   `json:"order_id"`
		Products  []*Product            `json:"products"`
		Status    enum.ProductionStatus `json:"status"`
		OrderDate time.Time             `json:"order_date"`
		CreatedAt time.Time             `json:"created_at"`
	}

	Product struct {
		ID          int    `json:"id"`
		ProductName string `json:"product_name"`
	}
)
