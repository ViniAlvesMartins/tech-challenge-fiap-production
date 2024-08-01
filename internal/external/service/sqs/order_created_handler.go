package sqs

import (
	"context"
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"log/slog"
	"time"
)

type (
	OrderCreatedMessage struct {
		ID          int              `json:"id" gorm:"primaryKey;autoIncrement"`
		OrderStatus enum.OrderStatus `json:"order_status"`
		OrderDate   time.Time        `json:"created_at"`
		Products    []*Product       `json:"products"`
	}

	Product struct {
		ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
		ProductName string `json:"product_name"`
	}

	OrderCreatedHandler struct {
		production contract.ProductionUseCase
		logger     *slog.Logger
	}
)

func NewOrderCreatedHandler(p contract.ProductionUseCase, l *slog.Logger) *OrderCreatedHandler {
	return &OrderCreatedHandler{production: p, logger: l}
}

func (f *OrderCreatedHandler) Handle(ctx context.Context, b []byte) error {
	var message OrderCreatedMessage

	f.logger.Info("Handling message...")

	if err := json.Unmarshal(b, &message); err != nil {
		return err
	}

	return nil
}
