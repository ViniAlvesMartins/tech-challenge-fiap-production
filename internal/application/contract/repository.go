package contract

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ProductionRepository interface {
	Create(ctx context.Context, production entity.Production) error
	GetByOrderId(ctx context.Context, orderId int) (*entity.Production, error)
	GetAll(ctx context.Context) ([]*entity.Production, error)
	UpdateStatusByOrderId(ctx context.Context, id int, status enum.ProductionStatus) error
}

type DynamoDB interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
}

type SnsService interface {
	SendMessage(ctx context.Context, message any) error
}
