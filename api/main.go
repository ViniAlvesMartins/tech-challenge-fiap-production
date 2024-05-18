package main

import (
	"context"
	"fmt"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/infra"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/use_case"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/database/dynamodb"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/handler/http_server"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/external/repository"
	"log/slog"
	"os"
)

// @title           Ze Burguer APIs
// @version         1.0
func main() {
	var err error
	var ctx = context.Background()
	var logger = loadLogger()

	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := dynamodb.NewConnection(cfg)

	if err != nil {
		logger.Error("error connecting tdo database", err)
		panic(err)
	}

	fmt.Println(db)

	productionRepository := repository.NewProductionRepository(db, logger)
	productionUseCase := use_case.NewPaymentUseCase(productionRepository, logger)

	app := http_server.NewApp(productionUseCase, logger)

	err = app.Run(ctx)

	if err != nil {
		logger.Error("error running application", err)
		panic(err)
	}
}

func loadConfig() (infra.Config, error) {
	return infra.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
