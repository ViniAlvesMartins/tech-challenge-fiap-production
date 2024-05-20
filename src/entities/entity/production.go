package entity

import (
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
)

type Production struct {
	OrderId      string                `json:"orderId"`
	ProductionId string                `json:"productionId"`
	CurrentState enum.ProductionStatus `json:"status"`
	CreatedAt    string                `json:"created_at,omitempty"`
}

func (p *Production) GetJSONValue() (string, error) {
	b, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(b), nil
}
