package output

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"time"
)

type ProductionDto struct {
	ID        int                   `json:"id"`
	OrderId   *int                  `json:"order_id"`
	Status    enum.StatusProduction `json:"status"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
}

func ProductionFromEntity(production entity.Production) ProductionDto {
	return ProductionDto{
		ID:        production.ID,
		OrderId:   production.OrderId,
		Status:    production.Status,
		CreatedAt: production.CreatedAt,
	}
}
