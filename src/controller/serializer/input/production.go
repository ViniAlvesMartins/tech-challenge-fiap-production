package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
)

type ProductionDto struct {
	OrderId      string                `json:"orderId"`
	ProductionId string                `json:"productionId"`
	CurrentState enum.ProductionStatus `json:"status"`
	CreatedAt    string                `json:"created_at,omitempty"`
}

func (p *ProductionDto) ConvertToEntity() entity.Production {
	return entity.Production{
		ProductionId: p.ProductionId,
		OrderId:      p.OrderId,
		CurrentState: p.CurrentState,
		CreatedAt:    p.CreatedAt,
	}
}
