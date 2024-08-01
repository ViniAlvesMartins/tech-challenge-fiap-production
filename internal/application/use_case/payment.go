package use_case

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

type PaymentUseCase struct {
	snsService contract.SnsService
}

func NewPaymentUseCase(s contract.SnsService) *PaymentUseCase {
	return &PaymentUseCase{
		snsService: s,
	}
}

func (p *PaymentUseCase) CancelPayment(ctx context.Context, orderId int) error {
	message := entity.PaymentStatusUpdatedMessage{
		OrderId: orderId,
		Status:  enum.PaymentStatusCanceled,
	}

	return p.snsService.SendMessage(ctx, message)
}
