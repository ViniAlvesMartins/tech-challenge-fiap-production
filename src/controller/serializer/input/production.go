package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"time"
)

type ProductionDto struct {
	OrderId      int                   `json:"orderId"`
	ProductionId string                `json:"productionId"`
	Status       enum.ProductionStatus `json:"status"`
	CreatedAt    time.Time             `json:"created_at,omitempty"`
}

func (p *ProductionDto) ConvertToEntity() entity.Production {
	return entity.Production{
		ProductionId: p.ProductionId,
		OrderId:      p.OrderId,
		Status:       p.Status,
		CreatedAt:    p.CreatedAt,
	}
}
