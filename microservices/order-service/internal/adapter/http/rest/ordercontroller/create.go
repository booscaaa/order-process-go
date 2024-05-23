package ordercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaa/order-process-go/microservices/order-service/internal/core/dto"
)

// Create implements domain.OrderController.
func (controller *controller) Create(response http.ResponseWriter, request *http.Request) {
	var orderInput dto.OrderInput

	err := json.NewDecoder(request.Body).Decode(&orderInput)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := controller.usecase.Create(request.Context(), orderInput)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(response).Encode(&order)
}
