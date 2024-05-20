package contract

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type QueueService interface {
	ReceiveMessage(queueUrl string) (*[]types.Message, error)
	DeleteMessage(queueURL string, receiptHandle string) error
	SendMessage(queueUrl string, message string, messageGroupId string) error
}
