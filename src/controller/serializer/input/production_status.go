package input

type StatusProductionDto struct {
	Status string `json:"status" validate:"required"`
}
