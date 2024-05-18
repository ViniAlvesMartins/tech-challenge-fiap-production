package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
	"time"
)

type ProductionDto struct {
	ID        int                   `json:"id"`
	OrderId   *int                  `json:"order_id"`
	Status    enum.StatusProduction `json:"status"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
}

func (p *ProductionDto) ConvertToEntity() entity.Production {
	return entity.Production{
		ID:        p.ID,
		OrderId:   p.OrderId,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
	}
}
