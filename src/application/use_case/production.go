package use_case

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"log/slog"
)

type ProductionUseCase struct {
	repository contract.ProductionRepository
	logger     *slog.Logger
}

func NewPaymentUseCase(r contract.ProductionRepository, logger *slog.Logger) *ProductionUseCase {
	return &ProductionUseCase{
		repository: r,
		logger:     logger,
	}
}

func (p *ProductionUseCase) UpdateStatusById(id int, status enum.ProductionStatus) (bool, error) {
	_, err := p.repository.UpdateStatusById(id, status)
	if err != nil {
		return false, err
	}

	//evento pedido atualizado
	return true, nil
}

func (p *ProductionUseCase) GetById(id int) (*entity.Production, error) {
	prodution, err := p.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return prodution, nil
}

func (p *ProductionUseCase) Create(production entity.Production) (*entity.Production, error) {
	production.Status = enum.AWAITING_PAYMENT

	productionNew, err := p.repository.Create(production)

	if err != nil {
		return nil, err
	}

	return productionNew, nil
}
