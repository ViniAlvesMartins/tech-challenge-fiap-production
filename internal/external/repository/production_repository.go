package repository

import (
	"context"
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/uuid"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/entity"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"sort"
	"time"
)

const table = "productions"

type ProductionRepository struct {
	db   contract.DynamoDB
	uuid uuid.Interface
}

func NewProductionRepository(db *dynamodb.Client, u uuid.Interface) *ProductionRepository {
	return &ProductionRepository{
		db:   db,
		uuid: u,
	}
}

func (p *ProductionRepository) Create(ctx context.Context, production entity.Production) error {
	production.ID = p.uuid.NewString()
	production.CreatedAt = time.Now()

	i, err := attributevalue.Marshal(production)
	if err != nil {
		return err
	}

	items := i.(*types.AttributeValueMemberM).Value
	input := &dynamodb.PutItemInput{
		Item:      items,
		TableName: aws.String(table),
	}

	if _, err = p.db.PutItem(ctx, input); err != nil {
		return err
	}

	return nil
}

func (p *ProductionRepository) GetById(ctx context.Context, id string) (*entity.Production, error) {
	var notFoundErr *types.ResourceNotFoundException
	var production *entity.Production

	out, err := p.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: id},
		},
	})

	if err != nil {
		if errors.As(err, &notFoundErr) {
			return nil, nil
		}

		return nil, err
	}

	if err = attributevalue.UnmarshalMap(out.Item, &production); err != nil {
		return nil, err
	}

	return production, nil
}

func (p *ProductionRepository) GetAll(ctx context.Context) ([]*entity.Production, error) {
	var productions []*entity.Production

	projection := expression.NamesList(
		expression.Name("id"),
		expression.Name("order_id"),
		expression.Name("products"),
		expression.Name("status"),
		expression.Name("order_date"),
		expression.Name("created_at"),
	)

	expr, err := expression.NewBuilder().
		WithFilter(expression.Name("status").NotEqual(expression.Value(enum.ProductionStatusFinished))).
		WithProjection(projection).
		Build()

	if err != nil {
		return nil, err
	}

	out, err := p.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(table),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
	})

	if err != nil {
		return nil, err
	}

	if err = attributevalue.UnmarshalListOfMaps(out.Items, &productions); err != nil {
		return nil, err
	}

	sort.SliceStable(productions, func(i, j int) bool {
		return productions[j].CreatedAt.After(productions[i].CreatedAt)
	})

	return productions, nil
}

func (p *ProductionRepository) UpdateStatusById(ctx context.Context, id string, status enum.ProductionStatus) error {
	_, err := p.db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("set status = :status"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":status": &types.AttributeValueMemberS{Value: string(status)},
		},
	})

	return err
}
