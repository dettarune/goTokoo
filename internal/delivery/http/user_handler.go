package http

import (
	"encoding/json"
	"net/http"

	"github.com/dettarune/goTokoo/internal/model"
	"github.com/dettarune/goTokoo/internal/usecase"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCaseDeps
}

func NewUserController(useCase *usecase.UserUseCaseDeps, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	request := new(model.RegsisterRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call usecase
	response, err := c.UseCase.Register(r.Context(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user: %+v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.RegisterResponse]{
		Data: response,
	}); err != nil {
		c.Log.Warnf("Failed to encode response: %+v", err)
	}
}
