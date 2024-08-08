package main

import (
	"context"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/dynamodb"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/sns"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-common/uuid"
	usecase "github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/config"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/external/handler/http_server"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/external/repository"
	snsproducer "github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/external/service/sns"
	"log/slog"
	"os"
)

func main() {
	var err error
	var ctx = context.Background()
	var logger = loadLogger()

	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := dynamodb.NewConnection(ctx)

	if err != nil {
		logger.Error("error connecting tdo database", err)
		panic(err)
	}

	orderSnsConnection, err := sns.NewConnection(ctx, cfg.OrderStatusUpdatedTopic)
	if err != nil {
		logger.Error("error connecting to sns", err)
		panic(err)
	}

	orderUseCase := usecase.NewOrderUseCase(snsproducer.NewService(orderSnsConnection))

	paymentSnsConnection, err := sns.NewConnection(ctx, cfg.ProductionFailedTopic)
	if err != nil {
		logger.Error("error connecting to sns", err)
		panic(err)
	}

	paymentUseCase := usecase.NewPaymentUseCase(snsproducer.NewService(paymentSnsConnection))

	productionRepository := repository.NewProductionRepository(db, loadUUID())
	productionUseCase := usecase.NewProductionUseCase(productionRepository, orderUseCase, paymentUseCase)

	app := http_server.NewApp(productionUseCase, logger)

	app.Run()
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
