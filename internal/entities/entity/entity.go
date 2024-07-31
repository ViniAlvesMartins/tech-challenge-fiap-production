package entity

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"time"
)

type Production struct {
	OrderId      int                   `json:"order_id"`
	ProductionId int                   `json:"production_id"`
	Status       enum.ProductionStatus `json:"status"`
	CreatedAt    time.Time             `json:"created_at,omitempty"`
}
