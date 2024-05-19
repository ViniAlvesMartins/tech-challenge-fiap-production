package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

type SnsService struct{}

func NewSnsService() *SnsService { return &SnsService{} }

func NewConnectionSns() (*sns.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)

	if err != nil {
		panic(err)
	}

	client := sns.NewFromConfig(cfg)

	return client, nil
}

func (s *SnsService) SendMessage(paymentId int, status enum.ProductionStatus) (bool, error) {

	client, _ := NewConnectionSns()

	message := &Message{
		OrderId: paymentId,
		Status:  status,
	}

	messageJs, _ := json.Marshal(message)

	snsMessage := string(messageJs)

	input := &sns.PublishInput{
		Message:  aws.String(snsMessage),
		TopicArn: aws.String("arn:aws:sns:us-east-1:682279319757:update_order_status-topic"),
	}

	result, err := client.Publish(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to publish message, %v", err)
	}

	fmt.Printf("Message ID: %s\n", *result.MessageId)

	return true, nil
}

type Message struct {
	OrderId int
	Status  enum.ProductionStatus
}
