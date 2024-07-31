package sqs

import (
	"context"
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"log/slog"
)

type (
	PaymentStatusUpdateMessage struct {
		OrderId int    `json:"order_id"`
		Status  string `json:"status"`
	}

	PaymentStatusUpdateHandler struct {
		payment contract.PaymentUseCase
		logger  *slog.Logger
	}
)

func NewPaymentStatusUpdateHandler(p contract.PaymentUseCase, l *slog.Logger) *PaymentStatusUpdateHandler {
	return &PaymentStatusUpdateHandler{payment: p, logger: l}
}

func (f *PaymentStatusUpdateHandler) Handle(ctx context.Context, b []byte) error {
	var message PaymentStatusUpdateMessage

	f.logger.Info("Handling message...")

	if err := json.Unmarshal(b, &message); err != nil {
		return err
	}

	return nil
}
