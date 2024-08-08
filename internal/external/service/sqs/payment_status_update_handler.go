package sqs

import (
	"context"
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"log/slog"
)

type (
	PaymentStatusUpdateMessage struct {
		OrderId int                `json:"order_id"`
		Status  enum.PaymentStatus `json:"status"`
	}

	PaymentStatusUpdateHandler struct {
		production contract.ProductionUseCase
		logger     *slog.Logger
	}
)

func NewPaymentStatusUpdateHandler(p contract.ProductionUseCase, l *slog.Logger) *PaymentStatusUpdateHandler {
	return &PaymentStatusUpdateHandler{production: p, logger: l}
}

func (f *PaymentStatusUpdateHandler) Handle(ctx context.Context, b []byte) error {
	var message PaymentStatusUpdateMessage

	f.logger.Info("Handling message...")

	if err := json.Unmarshal(b, &message); err != nil {
		return err
	}

	if message.Status != enum.PaymentStatusConfirmed {
		return nil
	}

	production, err := f.production.GetByOrderId(ctx, message.OrderId)
	if err != nil {
		return err
	}

	if production == nil {
		return nil
	}

	return f.production.UpdateStatusByOrderId(ctx, production.OrderId, enum.ProductionStatusPreparing)
}
