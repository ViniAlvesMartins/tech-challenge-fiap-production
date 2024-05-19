package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
)

type Production struct {
	OrderId      int                   `json:"orderId"`
	ProductionId string                `json:"productionId"`
	Status       enum.ProductionStatus `json:"status"`
	CreatedAt    time.Time             `json:"created_at,omitempty"`
}

func (p *Production) GetJSONValue() (string, error) {
	b, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(b), nil
}
