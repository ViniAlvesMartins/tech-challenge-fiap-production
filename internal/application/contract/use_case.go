package contract

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

type (
	ProductionUseCase interface {
		UpdateStatusById(ctx context.Context, id string, status enum.ProductionStatus) error
		GetById(ctx context.Context, id string) (*entity.Production, error)
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
