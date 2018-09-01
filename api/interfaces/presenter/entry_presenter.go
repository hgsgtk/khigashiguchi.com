package presenter

import "github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"

// GetEntriesResponse represents the format response.
type GetEntriesResponse struct {
	Entities []entity.Entry `json:"entities"`
}
