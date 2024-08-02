package http_server

import (
	"context"
	"errors"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/doc/swagger"
	_ "github.com/ViniAlvesMartins/tech-challenge-fiap-production/doc/swagger"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/controller/controller"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger/v2"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

func (e *App) Run() {
	router := mux.NewRouter()

	productionController := controller.NewProductionController(e.productionUseCase, e.logger)
	router.HandleFunc("/productions/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}", productionController.UpdateStatusById).Methods("PATCH")
	router.HandleFunc("/productions", productionController.GetAll).Methods("GET")

	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

	swagger.SwaggerInfo.Title = "Ze Burguer Productions API"
	swagger.SwaggerInfo.Version = "1.0"

	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Fatal(err)
	}
}
