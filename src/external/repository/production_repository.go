package repository

import (
	"context"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"log/slog"
)

type ProductionRepository struct {
	db     *dynamodb.Client
	logger *slog.Logger
}

func NewProductionRepository(db *dynamodb.Client, logger *slog.Logger) *ProductionRepository {
	return &ProductionRepository{
		db:     db,
		logger: logger,
	}
}

type Item struct {
	Id        string `json:"pk"`
	OrderId   string `json:"sk"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func (p *ProductionRepository) Create(production entity.Production) (*entity.Production, error) {

	table := "Productions"
	id := uuid.New().String()

	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"orderId":      &types.AttributeValueMemberN{Value: fmt.Sprint(production.OrderId)},
			"productionId": &types.AttributeValueMemberS{Value: id},
			"status":       &types.AttributeValueMemberS{Value: string(production.Status)},
			"createdAt":    &types.AttributeValueMemberS{Value: production.CreatedAt.String()},
		},
		TableName: aws.String(table),
	}

	_, err := p.db.PutItem(context.TODO(), input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	production.ProductionId = id

	return &production, nil
}

func (p *ProductionRepository) GetAll() ([]entity.Production, error) {

	production := &entity.Production{}

	table := "Productions"

	out, err := p.db.G(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"orderId": &types.AttributeValueMemberN{Value: fmt.Sprint(orderId)},
		},
	})

	if err != nil {
		panic(err)
	}

	err = attributevalue.UnmarshalMap(out.Item, &production)

	if err != nil {
		panic(err)
	}

	return production, nil
}

func (p *ProductionRepository) GetById(productionId int) (*entity.Production, error) {

	production := &entity.Production{}

	table := "Payments"

	out, err := p.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"orderId": &types.AttributeValueMemberN{Value: fmt.Sprint(productionId)},
		},
	})

	if err != nil {
		panic(err)
	}

	err = attributevalue.UnmarshalMap(out.Item, &production)

	if err != nil {
		panic(err)
	}

	return production, nil
}

func (p *ProductionRepository) UpdateStatusById(paymentId int, status enum.ProductionStatus) (bool, error) {

	table := "Productions"

	_, err := p.db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"paymentId": &types.AttributeValueMemberN{Value: fmt.Sprint(paymentId)},
		},
		UpdateExpression: aws.String("set status = :status"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":status": &types.AttributeValueMemberS{Value: string(status)},
		},
	})

	if err != nil {
		panic(err)
	}

	return true, nil
}
