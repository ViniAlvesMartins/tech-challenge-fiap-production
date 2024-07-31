package repository

import (
	"context"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"log/slog"
	"sort"
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

func (p *ProductionRepository) Create(production entity.Production) (*entity.Production, error) {
	table := "productions"
	id := uuid.New().String()

	input := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"orderId":      &types.AttributeValueMemberS{Value: production.OrderId},
			"productionId": &types.AttributeValueMemberS{Value: id},
			"currentState": &types.AttributeValueMemberS{Value: string(production.CurrentState)},
			"createdAt":    &types.AttributeValueMemberS{Value: time.Now().Format(time.DateTime)},
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

func (p *ProductionRepository) GetAll() ([]*entity.Production, error) {

	var productions []*entity.Production

	table := "productions"

	cond1 := expression.Name("currentState").NotEqual(expression.Value("FINISHED"))
	proj := expression.NamesList(
		expression.Name("orderId"),
		expression.Name("productionId"),
		expression.Name("currentState"),
		expression.Name("createdAt"),
	)
	expr, err := expression.NewBuilder().
		WithFilter(cond1).
		WithProjection(proj).
		Build()
	if err != nil {
		fmt.Println(err)
	}

	out, err := p.db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String(table),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
	})

	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &productions)

	if err != nil {
		return nil, err
	}

	sort.SliceStable(productions, func(i, j int) bool {
		return productions[i].CreatedAt < productions[j].CreatedAt
	})

	return productions, nil
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
