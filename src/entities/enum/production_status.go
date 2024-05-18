package enum

import "slices"

type ProductionStatus string

const (
	AWAITING_PAYMENT ProductionStatus = "AWAITING_PAYMENT"
	RECEIVED         ProductionStatus = "RECEIVED"
	PREPARING        ProductionStatus = "PREPARING"
	READY            ProductionStatus = "READY"
	FINISHED         ProductionStatus = "FINISHED"
)

func ValidateStatus(val string) bool {
	validStatus := []ProductionStatus{AWAITING_PAYMENT, RECEIVED, PREPARING, READY, FINISHED}
	return slices.Contains(validStatus, ProductionStatus(val))
}
