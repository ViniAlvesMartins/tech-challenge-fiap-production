package enum

import "slices"

type ProductionStatus string

const (
	PREPARING ProductionStatus = "PREPARING"
	READY     ProductionStatus = "READY"
	FINISHED  ProductionStatus = "FINISHED"
)

func ValidateStatus(val string) bool {
	validStatus := []ProductionStatus{PREPARING, READY, FINISHED}
	return slices.Contains(validStatus, ProductionStatus(val))
}
