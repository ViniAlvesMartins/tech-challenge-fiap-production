package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/controller/serializer/output"
	"github.com/ViniAlvesMartins/tech-challenge-fiap/src/entities/enum"
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

// CreateProduction godoc
// @Summary      Create production
// @Description  Place a new production
// @Tags         Productions
// @Accept       json
// @Produce      json
// @Param        request   body      input.ProductionDto  true  "Production properties"
// @Success      201  {object}  Response{data=output.ProductionDto}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Failure      404  {object}  swagger.ResourceNotFoundResponse{data=interface{}}
// @Router       /productions [post]
func (p *ProductionController) CreateProduction(w http.ResponseWriter, r *http.Request) {
	var productionDto input.ProductionDto

	if err := json.NewDecoder(r.Body).Decode(&productionDto); err != nil {
		p.logger.Error("unable to decode the request body", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(
			Response{
				Error: "Unable to decode the request body",
				Data:  nil,
			})
		if err != nil {
			return
		}
	}

	production, err := p.productionUseCase.Create(productionDto.ConvertToEntity())
	if err != nil {
		p.logger.Error("error creating order", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(
			Response{
				Error: "Error creating order",
				Data:  nil,
			})
		if err != nil {
			return
		}
		return
	}

	orderOutput := output.ProductionFromEntity(production)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(
		Response{
			Error: "",
			Data:  orderOutput,
		})
	if err != nil {
		return
	}
}

// UpdateProductionStatusById godoc
// @Summary      Find production
// @Description  Find production by id
// @Tags         Production
// @Produce      json
// @Param        id   path      int  true  "Production ID"
// @Param        request   body      input.StatusProductionDto  true  "Production status"
// @Success      204  {object}  interface{}
// @Failure      500  {object}  swagger.InternalServerErrorResponse{data=interface{}}
// @Failure      404  {object}  swagger.ResourceNotFoundResponse{data=interface{}}
// @Router       /productions/{id} [patch]
func (p *ProductionController) UpdateProductionStatusById(w http.ResponseWriter, r *http.Request) {
	productionIdParam := mux.Vars(r)["productionId"]
	productionId, err := strconv.Atoi(productionIdParam)

	if err != nil {
		p.logger.Error("error converting productionId to int", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(Response{
			Error: "Error to convert productionId to int",
			Data:  nil,
		})
		if err != nil {
			return
		}
	}

	var statusProductionDto input.StatusProductionDto
	if err := json.NewDecoder(r.Body).Decode(&statusProductionDto); err != nil {
		p.logger.Error("unable to decode the request body", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(
			Response{
				Error: "Unable to decode the request body",
				Data:  nil,
			})
		if err != nil {
			return
		}
	}

	if !enum.ValidateStatus(statusProductionDto.Status) {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(Response{
			Error: "Invalid status",
			Data:  nil,
		})
		if err != nil {
			return
		}
	}

	production, err := p.productionUseCase.GetById(productionId)
	if err != nil {
		p.logger.Error("error getting production by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(Response{
			Error: "Error to get production",
			Data:  nil,
		})
		if err != nil {
			return
		}
	}

	if production == nil {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(
			Response{
				Error: "production not found",
				Data:  nil,
			})
		if err != nil {
			return
		}
	}

	if err := p.productionUseCase.UpdateStatusById(productionId, enum.StatusProduction(statusProductionDto.Status)); err != nil {
		p.logger.Error("error updating status by id", slog.Any("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(Response{
			Error: "Error updating status",
			Data:  nil,
		})
		if err != nil {
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
