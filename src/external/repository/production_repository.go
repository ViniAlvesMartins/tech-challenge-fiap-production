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
	"strconv"
	"time"
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

	table := "productions"
	id := uuid.New().String()

	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"orderId":      &types.AttributeValueMemberS{Value: production.OrderId},
			"productionId": &types.AttributeValueMemberS{Value: id},
			"currentState": &types.AttributeValueMemberS{Value: string(production.CurrentState)},
			"createdAt":    &types.AttributeValueMemberS{Value: time.Now().String()},
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

func (p *ProductionRepository) GetById(orderId int) (*entity.Production, error) {

	production := &entity.Production{}

	table := "productions"

	out, err := p.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"orderId": &types.AttributeValueMemberS{Value: strconv.Itoa(orderId)},
		},
	})

	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(out.Item, &production)

	if err != nil {
		return nil, err
	}

	return production, nil
}

func (p *ProductionRepository) UpdateStatusById(orderId int, status enum.ProductionStatus) (bool, error) {
	table := "productions"

	_, err := p.db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"orderId": &types.AttributeValueMemberS{Value: strconv.Itoa(orderId)},
		},
		UpdateExpression: aws.String("set currentState = :currentState"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":currentState": &types.AttributeValueMemberS{Value: string(status)},
		},
	})

	if err != nil {
		panic(err)
	}

	return true, nil
}
