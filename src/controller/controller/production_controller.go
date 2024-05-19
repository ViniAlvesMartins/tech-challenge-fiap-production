package controller

import (
	"encoding/json"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/application/contract"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/controller/serializer/input"
	"github.com/ViniAlvesMartins/tech-challenge-fiap-production/src/entities/enum"
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
	orderIdParam := mux.Vars(r)["orderId"]
	orderId, err := strconv.Atoi(orderIdParam)

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

	production, err := p.productionUseCase.GetById(orderId)
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

	_, err = p.productionUseCase.UpdateStatusById(orderId, enum.ProductionStatus(statusProductionDto.Status))
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
