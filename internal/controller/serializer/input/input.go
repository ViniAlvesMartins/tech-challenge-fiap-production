package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

type (
	ProductionDto struct {
		OrderId      string                `json:"orderId"`
		ProductionId string                `json:"productionId"`
		CurrentState enum.ProductionStatus `json:"status"`
		CreatedAt    string                `json:"created_at,omitempty"`
	}

	StatusProductionDto struct {
		Status string `json:"status" validate:"required"`
	}
)

func (p *ProductionDto) ConvertToEntity() entity.Production {
	return entity.Production{
		ProductionId: p.ProductionId,
		OrderId:      p.OrderId,
		CurrentState: p.CurrentState,
		CreatedAt:    p.CreatedAt,
	}
}
