package use_case

import (
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"log/slog"
)

type ProductionUseCase struct {
	repository contract.ProductionRepository
	snsService contract.SnsService
	logger     *slog.Logger
}

func NewPaymentUseCase(r contract.ProductionRepository, s contract.SnsService, logger *slog.Logger) *ProductionUseCase {
	return &ProductionUseCase{
		repository: r,
		snsService: s,
		logger:     logger,
	}
}

func (p *ProductionUseCase) UpdateStatusById(id int, status enum.ProductionStatus) (bool, error) {
	_, err := p.repository.UpdateStatusById(id, status)
	if err != nil {
		return false, err
	}

	res, err := p.snsService.SendMessage(id, status)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	return true, nil
}

func (p *ProductionUseCase) GetById(id int) (*entity.Production, error) {
	prodution, err := p.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return prodution, nil
}

func (p *ProductionUseCase) GetAll() ([]*entity.Production, error) {
	productions, err := p.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return productions, nil
}

func (p *ProductionUseCase) Create(production entity.Production) (*entity.Production, error) {
	productionNew, err := p.repository.Create(production)

	if err != nil {
		return nil, err
	}

	return productionNew, nil
}
