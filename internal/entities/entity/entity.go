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
		ID        string                `json:"id" dynamodbav:"id"`
		OrderId   int                   `json:"order_id" dynamodbav:"order_id"`
		Products  []*Product            `json:"products" dynamodbav:"products"`
		Status    enum.ProductionStatus `json:"status" dynamodbav:"status"`
		OrderDate time.Time             `json:"order_date" dynamodbav:"order_date"`
		CreatedAt time.Time             `json:"created_at" dynamodbav:"created_at"`
	}

	Product struct {
		ID          int    `json:"id" dynamodbav:"id"`
		ProductName string `json:"product_name" dynamodbav:"product_name"`
	}
)
