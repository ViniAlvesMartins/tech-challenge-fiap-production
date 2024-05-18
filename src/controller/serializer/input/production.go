package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"time"
)

type ProductionDto struct {
	ID        string                `json:"id"`
	OrderId   *int                  `json:"order_id"`
	Status    enum.ProductionStatus `json:"status"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
}

func (p *ProductionDto) ConvertToEntity() entity.Production {
	return entity.Production{
		ProductionId: p.ID,
		OrderId:      p.OrderId,
		Status:       p.Status,
		CreatedAt:    p.CreatedAt,
	}
}
