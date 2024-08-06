package contract

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

type (
	ProductionUseCase interface {
		UpdateStatusByOrderId(ctx context.Context, orderId int, status enum.ProductionStatus) error
		GetByOrderId(ctx context.Context, orderId int) (*entity.Production, error)
		GetAll(ctx context.Context) ([]*entity.Production, error)
		Create(ctx context.Context, production entity.Production) error
	}

	OrderUseCase interface {
		UpdateOrderStatus(ctx context.Context, orderId int, status enum.ProductionStatus) error
	}

	PaymentUseCase interface {
		CancelPayment(ctx context.Context, orderId int) error
	}
)
