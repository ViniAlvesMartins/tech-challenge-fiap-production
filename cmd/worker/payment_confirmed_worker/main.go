package main

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/dynamodb"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/sns"
	sqsservice "github.com/ViniAlvesMartins/tech-challenge-fiap-common/sqs"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/uuid"
	usecase "github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/config"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/external/repository"
	snsproducer "github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/external/service/sns"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/external/service/sqs"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var err error
	var ctx, cancel = context.WithCancel(context.Background())
	var logger = loadLogger()

	logger.Info("Initializing worker...")
	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := dynamodb.NewConnection(ctx)
	if err != nil {
		logger.Error("error connecting to database", err)
		panic(err)
	}

	consumer, err := sqsservice.NewConnection(ctx, cfg.PaymentStatusUpdatedQueue, 1, 20)
	if err != nil {
		logger.Error("error connecting to sqs", err)
		panic(err)
	}

	orderSnsConnection, err := sns.NewConnection(ctx, cfg.OrderStatusUpdatedTopic)
	if err != nil {
		logger.Error("error connecting to order sns", err)
		panic(err)
	}

	paymentSnsConnection, err := sns.NewConnection(ctx, cfg.PaymentStatusUpdatedTopic)
	if err != nil {
		logger.Error("error connecting to payment sns", err)
		panic(err)
	}

	orderSnsService := snsproducer.NewService(orderSnsConnection)
	paymentSnsService := snsproducer.NewService(paymentSnsConnection)

	orderUseCase := usecase.NewOrderUseCase(orderSnsService)
	paymentUseCase := usecase.NewPaymentUseCase(paymentSnsService)

	productionRepository := repository.NewProductionRepository(db, loadUUID())

	productionUseCase := usecase.NewProductionUseCase(productionRepository, orderUseCase, paymentUseCase)
	paymentUpdatedHandler := sqs.NewPaymentStatusUpdateHandler(productionUseCase, logger)

	logger.Info("Starting consumer...")
	paymentConfirmedConsumer := sqs.NewConsumer(consumer, paymentUpdatedHandler, logger)

	var wg sync.WaitGroup
	wg.Add(1)
	go paymentConfirmedConsumer.Start(ctx, &wg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	cancel()
	wg.Wait()
	logger.Info("Finishing worker...")
}

func loadUUID() uuid.Interface {
	return &uuid.UUID{}
}

func loadConfig() (config.Config, error) {
	return config.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
