package output

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"time"
)

type ProductionDto struct {
	ID        int                   `json:"id"`
	OrderId   int                   `json:"order_id"`
	Status    enum.ProductionStatus `json:"status"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
}
