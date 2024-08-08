package sns

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/sns"
)

type Service struct {
	sns *sns.Service
}

func NewService(sns *sns.Service) *Service {
	return &Service{
		sns: sns,
	}
}

func (s *Service) SendMessage(ctx context.Context, message any) error {
	return s.sns.Publish(ctx, message)
}
