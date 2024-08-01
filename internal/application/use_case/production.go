package use_case

import (
	"context"
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

var ErrItemNotFound = errors.New("item not found in production")

type ProductionUseCase struct {
	repository     contract.ProductionRepository
	orderUseCase   contract.OrderUseCase
	paymentUseCase contract.PaymentUseCase
}

func NewProductionUseCase(r contract.ProductionRepository, o contract.OrderUseCase, p contract.PaymentUseCase) *ProductionUseCase {
	return &ProductionUseCase{
		repository:     r,
		orderUseCase:   o,
		paymentUseCase: p,
	}
}

func (p *ProductionUseCase) UpdateStatusById(ctx context.Context, id string, status enum.ProductionStatus) error {
	production, err := p.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if production == nil {
		return ErrItemNotFound
	}

	err = p.repository.UpdateStatusById(ctx, id, status)
	if err != nil {
		return err
	}

	if status == enum.ProductionStatusCanceled {
		return p.paymentUseCase.CancelPayment(ctx, production.OrderId)
	}

	return p.orderUseCase.UpdateOrderStatus(ctx, production.OrderId, status)
}

func (p *ProductionUseCase) GetById(ctx context.Context, id string) (*entity.Production, error) {
	return p.repository.GetById(ctx, id)
}

func (p *ProductionUseCase) GetAll(ctx context.Context) ([]*entity.Production, error) {
	return p.repository.GetAll(ctx)
}

func (p *ProductionUseCase) Create(ctx context.Context, production entity.Production) error {
	return p.repository.Create(ctx, production)
}
