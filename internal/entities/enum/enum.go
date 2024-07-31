package enum

import "slices"

type (
	ProductionStatus string
	OrderStatus      string
)

const (
	ProductionStatusPreparing ProductionStatus = "PREPARING"
	ProductionStatusReady     ProductionStatus = "READY"
	ProductionStatusFinished  ProductionStatus = "FINISHED"

	OrderStatusAwaitingPayment OrderStatus = "AWAITING_PAYMENT"
	OrderStatusReceived        OrderStatus = "RECEIVED"
	OrderStatusPreparing       OrderStatus = "PREPARING"
	OrderStatusReady           OrderStatus = "READY"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusFinished        OrderStatus = "FINISHED"
)

func ValidateStatus(val string) bool {
	validStatus := []ProductionStatus{ProductionStatusPreparing, ProductionStatusReady, ProductionStatusFinished}
	return slices.Contains(validStatus, ProductionStatus(val))
}
