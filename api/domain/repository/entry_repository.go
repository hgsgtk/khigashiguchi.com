package repository

import "github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"

// EntryRepository is the interface which get all Entry entities.
type EntryRepository interface {
	GetAll() ([]entity.Entry, error)
}
