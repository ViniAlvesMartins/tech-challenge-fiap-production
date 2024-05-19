package sqs

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"

	"log"
	"log/slog"
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
			continue
		}

		if result == nil {
			log.Printf("Failed result sqs %v", result)
			continue
		} else {

			sqsMessage := &SqsMessage{}

			err := json.Unmarshal([]byte(*result.Body), &sqsMessage)
			if err != nil {
				return err
			}

			sqsMessageReturn := &SqsMessageReturn{}

			log.Println(sqsMessage.Message)

			err = json.Unmarshal([]byte(sqsMessage.Message), &sqsMessageReturn)
			if err != nil {
				return err
			}

			var production input.ProductionDto
			production.OrderId = sqsMessageReturn.OrderId
			production.Status = enum.PREPARING

			productionEntity := production.ConvertToEntity()

			_, err = s.productionUseCase.Create(productionEntity)
			if err != nil {
				return err
			}

			if err != nil {
				s.logger.Error("error create production", slog.Any("error", err.Error()))
			} else {
				log.Printf(*result.ReceiptHandle)
				err := s.sqsService.DeleteMessage(queueUrl, *result.ReceiptHandle)
				if err != nil {
					return err
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
