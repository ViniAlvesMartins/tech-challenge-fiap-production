package http_server

import (
	"context"
	_ "github.com/ViniAlvesMartins/tech-challenge-fiap-production/doc/swagger"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/controller/controller"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger/v2"
	"log/slog"
	"net/http"
)

type App struct {
	productionUseCase contract.ProductionUseCase
	logger            *slog.Logger
}

func NewApp(
	productionUseCase contract.ProductionUseCase,
	logger *slog.Logger,
) *App {
	return &App{
		logger:            logger,
		productionUseCase: productionUseCase,
	}
}

func (e *App) Run(ctx context.Context) error {
	router := mux.NewRouter()

	productionController := controller.NewProductionController(e.productionUseCase, e.logger)
	router.HandleFunc("/productions/{productionId:[0-9]+}", productionController.UpdateProductionStatusById).Methods("PATCH")
	router.HandleFunc("/productions", productionController.GetAllProductions).Methods("GET")

	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	return http.ListenAndServe(":8082", router)
}
