package enum

import (
	"errors"
	"slices"
)

var ErrInvalidProductionStatus = errors.New("invalid production status")

type (
	ProductionStatus string
	OrderStatus      string
	PaymentStatus    string
)

const (
	ProductionAwaitingPayment ProductionStatus = "AWAITING_PAYMENT"
	ProductionStatusReceived  ProductionStatus = "RECEIVED"
	ProductionStatusPreparing ProductionStatus = "PREPARING"
	ProductionStatusReady     ProductionStatus = "READY"
	ProductionStatusFinished  ProductionStatus = "FINISHED"
	ProductionStatusCanceled  ProductionStatus = "CANCELED"

	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusConfirmed PaymentStatus = "CONFIRMED"
	PaymentStatusCanceled  PaymentStatus = "CANCELED"

	OrderStatusReceived        OrderStatus = "RECEIVED"
	OrderStatusPreparing       OrderStatus = "PREPARING"
	OrderStatusReady           OrderStatus = "READY"
	OrderStatusAwaitingPayment OrderStatus = "AWAITING_PAYMENT"
	OrderStatusFinished        OrderStatus = "FINISHED"
)

func ValidateProductionStatus(val string) error {
	validStatus := []ProductionStatus{ProductionStatusCanceled, ProductionStatusReceived, ProductionStatusPreparing, ProductionStatusReady, ProductionStatusFinished}

	if !slices.Contains(validStatus, ProductionStatus(val)) {
		return ErrInvalidProductionStatus
	}

	return nil
}
