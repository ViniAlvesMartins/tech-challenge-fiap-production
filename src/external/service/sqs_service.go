package service

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SqsService struct{}

func NewSqsService() *SqsService { return &SqsService{} }

func NewConnectionSqs() (*sqs.Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)

	if err != nil {
		fmt.Println(err)
	}

	client := sqs.NewFromConfig(cfg)

	return client, nil
}

func (s *SqsService) ReceiveMessage(queueURL string) (*[]types.Message, error) {

	client, _ := NewConnectionSqs()

	out, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     20,
	})

	if err != nil {
		log.Printf("Failed to fetch sqs message %v", err)
	}

	if len(out.Messages) >= 1 {
		return &out.Messages, nil
	}

	return nil, nil
}

func (s *SqsService) DeleteMessage(queueURL string, receiptHandle string) error {

	client, _ := NewConnectionSqs()

	_, err := client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: &receiptHandle,
	})

	if err != nil {
		log.Printf("Delete Error %v", err)
	}

	return nil
}

func (s *SqsService) SendMessage(queueURL string, message string, messageGroupId string) error {

	client, _ := NewConnectionSqs()

	// Enviar a mensagem para a fila
	result, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:               &queueURL,
		MessageBody:            &message,
		MessageGroupId:         &messageGroupId,
		MessageDeduplicationId: &message,
	})

	if err != nil {
		fmt.Println("Error sending message to SQS:", err)
	}

	fmt.Println("Message ID:", *result.MessageId)

	return nil
}
