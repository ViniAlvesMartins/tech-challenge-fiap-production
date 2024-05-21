package sqs

import (
	"encoding/json"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
	"log"
	"log/slog"
	"strconv"
	"time"
)

type SqsConsumer struct {
	sqsService        contract.QueueService
	productionUseCase contract.ProductionUseCase
	logger            *slog.Logger
}

func NewSqsConsumer(queueService contract.QueueService,
	productionUseCase contract.ProductionUseCase,
	logger *slog.Logger) *SqsConsumer {
	return &SqsConsumer{
		sqsService:        queueService,
		productionUseCase: productionUseCase,
		logger:            logger,
	}
}

func (s *SqsConsumer) Run() error {
	queueUrl := "https://sqs.us-east-1.amazonaws.com/682279319757/to_production_order_queue.fifo"

	for {
		result, err := s.sqsService.ReceiveMessage(queueUrl)
		if err != nil {
			log.Printf("Failed to fetch sqs message %v", err)
		}
		if result == nil {
			log.Printf("Failed result sqs %v", result)
		} else {
			for _, message := range *result {
				fmt.Println("message: ", *message.Body)

				sqsMessageReturn := &SqsMessageReturn{}
				json.Unmarshal([]byte(*message.Body), &sqsMessageReturn)

				var production input.ProductionDto
				production.OrderId = strconv.Itoa(sqsMessageReturn.OrderId)
				production.CurrentState = enum.ProductionStatus(sqsMessageReturn.Status)

				_, err = s.productionUseCase.Create(production.ConvertToEntity())
				if err != nil {
					s.logger.Error("error create production", slog.Any("error", err.Error()))
					return err
				} else {
					log.Printf(*message.ReceiptHandle)
					err := s.sqsService.DeleteMessage(queueUrl, *message.ReceiptHandle)
					if err != nil {
						return err
					}
				}

			}
		}

		time.Sleep(5 * time.Second)
	}
}

type SqsMessage struct {
	Type             string
	MessageId        string
	TopicArn         string
	Message          string
	Timestamp        string
	SignatureVersion string
	Signature        string
	SigningCertURL   string
	UnsubscribeURL   string
}

type SqsMessageReturn struct {
	OrderId int
	Status  string
}
