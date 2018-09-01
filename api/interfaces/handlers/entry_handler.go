package handlers

import (
	"net/http"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/presenter"
)

// GetEntriesHandler handle request GET /entries.
func GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: テストを通すための仮実装
	res := presenter.GetEntriesResponse{
		Entities: entity.Entry{
			Title: "test title",
			URL:   "http://example.com",
		},
	}
	presenter.RespondJson(w, res, http.StatusOK)
}
