package handlers

import (
	"net/http"

	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/presenter"
	"github.com/Khigashiguchi/khigashiguchi.com/api/usecase"
)

// Handler is interface of handling request.
type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type getEntriesHandler struct {
	UseCase usecase.GetEntriesUseCase
}

// GetEntriesHandler handle request GET /entries.
func (h *getEntriesHandler) Handler(w http.ResponseWriter, r *http.Request) {
	entries, err := h.UseCase.Run()
	if err != nil {
		presenter.RespondErrorJson(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	res := presenter.GetEntriesResponse{
		Entities: entries,
	}
	presenter.RespondJson(w, res, http.StatusOK)
}

// NewGetEntriesHandler create new handler of getting entries.
func NewGetEntriesHandler(db repository.Executor) Handler {
	return &getEntriesHandler{UseCase: usecase.NewGetEntriesUseCase(db)}
}
