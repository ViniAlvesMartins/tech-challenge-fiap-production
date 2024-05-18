package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
)

type Production struct {
	ProductionId int                   `json:"production_id"`
	OrderId      *int                  `json:"order_id"`
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
