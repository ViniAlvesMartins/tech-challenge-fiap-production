package contract

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
)

type ProductionUseCase interface {
	UpdateStatusById(id int, status enum.ProductionStatus) error
	GetById(id int) (*entity.Production, error)
	GetAll() ([]entity.Production, error)
	Create(production entity.Production) (*entity.Production, error)
}
