package use_case

import (
	"context"
	"errors"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

var (
	ErrItemNotFound  = errors.New("item not found in production")
	ErrInvalidStatus = errors.New("production item cannot be created with this status")
)

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

func (p *ProductionUseCase) UpdateStatusByOrderId(ctx context.Context, orderId int, status enum.ProductionStatus) error {
	production, err := p.repository.GetByOrderId(ctx, orderId)
	if err != nil {
		return err
	}

	if production == nil {
		return ErrItemNotFound
	}

	err = p.repository.UpdateStatusByOrderId(ctx, orderId, status)
	if err != nil {
		return err
	}

	if status == enum.ProductionStatusCanceled {
		return p.paymentUseCase.CancelPayment(ctx, production.OrderId)
	}

	return p.orderUseCase.UpdateOrderStatus(ctx, production.OrderId, status)
}

func (p *ProductionUseCase) GetByOrderId(ctx context.Context, orderId int) (*entity.Production, error) {
	return p.repository.GetByOrderId(ctx, orderId)
}

func (p *ProductionUseCase) GetAll(ctx context.Context) ([]*entity.Production, error) {
	return p.repository.GetAll(ctx)
}

func (p *ProductionUseCase) Create(ctx context.Context, production entity.Production) error {
	if production.Status != enum.ProductionAwaitingPayment {
		return fmt.Errorf("%w - %s", ErrInvalidStatus, production.Status)
	}

	return p.repository.Create(ctx, production)
}
