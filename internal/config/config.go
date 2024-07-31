package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DynamoDBRegion       string `envconfig:"dynamodb_region"`
	DynamoDBUrl          string `envconfig:"dynamodb_url"`
	DynamoDBAccessKey    string `envconfig:"dynamodb_access_key"`
	DynamoDBSecretAccess string `envconfig:"dynamodb_secret_access"`

	SnsRegion                   string `envconfig:"sns_region"`
	SnsUrl                      string `envconfig:"sns_url"`
	SnsAccessKey                string `envconfig:"sns_access_key"`
	OrderStatusUpdatedQueue     string `envconfig:"order_status_updated_queue"`
	PaymentStatusUpdatedQueue   string `envconfig:"payment_status_updated_queue"`
	ProductionOrderCreatedQueue string `envconfig:"production_order_created_queue"`
	SnsSecretAccess             string `envconfig:"sns_secret_access"`
}

func NewConfig() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return
}
