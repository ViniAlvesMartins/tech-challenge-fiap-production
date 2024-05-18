package contract

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type ProductionRepository interface {
	UpdateStatusById(id int, status enum.ProductionStatus) (bool, error)
	GetById(id int) (*entity.Production, error)
	GetAll() (*[]entity.Production, error)
	Create(production entity.Production) (*entity.Production, error)
}
