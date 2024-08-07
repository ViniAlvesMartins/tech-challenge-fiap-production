package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DynamoDBRegion       string `envconfig:"dynamodb_region"`
	DynamoDBUrl          string `envconfig:"dynamodb_url"`
	DynamoDBAccessKey    string `envconfig:"dynamodb_access_key"`
	DynamoDBSecretAccess string `envconfig:"dynamodb_secret_access"`

	SnsRegion       string `envconfig:"sns_region"`
	SnsUrl          string `envconfig:"sns_url"`
	SnsAccessKey    string `envconfig:"sns_access_key"`
	SnsSecretAccess string `envconfig:"sns_secret_access"`

	OrderStatusUpdatedTopic string `envconfig:"order_status_updated_topic"`
	ProductionFailedTopic   string `envconfig:"production_failed_topic"`

	PaymentStatusUpdatedQueue   string `envconfig:"payment_status_updated_queue"`
	ProductionOrderCreatedQueue string `envconfig:"production_order_created_queue"`
}

func NewConfig() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return
}
