package use_case

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

type OrderUseCase struct {
	snsService contract.SnsService
}

func NewOrderUseCase(s contract.SnsService) *OrderUseCase {
	return &OrderUseCase{
		snsService: s,
	}
}

func (p *OrderUseCase) UpdateOrderStatus(ctx context.Context, orderId int, status enum.ProductionStatus) error {
	if err := enum.ValidateProductionStatus(string(status)); err != nil {
		return err
	}

	message := entity.OrderStatusUpdatedMessage{
		OrderId: orderId,
		Status:  enum.OrderStatus(status),
	}

	return p.snsService.SendMessage(ctx, message)
}
