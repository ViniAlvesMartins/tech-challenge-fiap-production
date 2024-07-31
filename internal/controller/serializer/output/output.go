package output

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"time"
)

type ProductionDto struct {
	ID        int                   `json:"id"`
	OrderId   int                   `json:"order_id"`
	Status    enum.ProductionStatus `json:"status"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
}

func ProductionFromEntity(production *entity.Production) ProductionDto {
	return ProductionDto{
		ID:        production.ProductionId,
		OrderId:   production.OrderId,
		Status:    production.Status,
		CreatedAt: production.CreatedAt,
	}
}
