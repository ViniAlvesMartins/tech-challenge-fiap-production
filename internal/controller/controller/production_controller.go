package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/internal/entities/enum"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

type ProductionController struct {
	productionUseCase contract.ProductionUseCase
	logger            *slog.Logger
}

func NewProductionController(productionUseCase contract.ProductionUseCase, logger *slog.Logger) *ProductionController {
	return &ProductionController{
		productionUseCase: productionUseCase,
		logger:            logger,
	}
}

// GetAll godoc
// @Summary      Find all production
// @Description  Find all production by id
// @Tags         Production
// @Produce      json
// @Success      200  {object}  interface{}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Failure      404  {object}  swagger.ResourceNotFoundResponse{data=interface{}}
// @Router       /productions [get]
func (p *ProductionController) GetAll(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	productions, err := p.productionUseCase.GetAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		p.logger.Error("error listing productions: ", slog.Any("error", err.Error()))
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "error listing productions",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(
		Response{
			Error: "",
			Data:  productions,
		})
	w.Write(jsonResponse)
	return
}

// UpdateStatusById godoc
// @Summary      Update production
// @Description  Update production by id
// @Tags         Production
// @Produce      json
// @Param        id   path      int  true  "Production ID"
// @Param        request   body      input.StatusProductionDto  true  "Production status"
// @Success      204  {object}  interface{}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Failure      404  {object}  swagger.ResourceNotFoundResponse{data=interface{}}
// @Router       /productions/{id} [patch]
func (p *ProductionController) UpdateStatusById(w http.ResponseWriter, r *http.Request) {
	var statusProductionDto input.StatusProductionDto
	var id = mux.Vars(r)["id"]
	var ctx = r.Context()
	orderId, err := strconv.Atoi(id)

	if err != nil {
		p.logger.Error("error converting order id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "error converting order id",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&statusProductionDto); err != nil {
		p.logger.Error("unable to decode the request body", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "unable to decode request body",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	if err := enum.ValidateProductionStatus(statusProductionDto.Status); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "invalid status",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	production, err := p.productionUseCase.GetById(ctx, orderId)
	if err != nil {
		p.logger.Error("error getting production by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "error getting production",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	if production == nil {
		w.WriteHeader(http.StatusNotFound)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "error getting order not found",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	err = p.productionUseCase.UpdateStatusById(ctx, id, enum.ProductionStatus(statusProductionDto.Status))
	if err != nil {
		p.logger.Error("error updating status", slog.Any("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse, _ := json.Marshal(
			Response{
				Error: "error updating status",
				Data:  nil,
			})
		w.Write(jsonResponse)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
