package input

import (
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
)

type (
	ProductionDto struct {
		OrderId      string                `json:"order_id"`
		ProductionId string                `json:"production_id"`
		CurrentState enum.ProductionStatus `json:"status"`
		CreatedAt    string                `json:"created_at,omitempty"`
	}

	StatusProductionDto struct {
		Status string `json:"status" validate:"required"`
	}
)
