package usecase

import (
	"fmt"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/logger"
	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
)

// GetEntriesUseCaseImpl represents the interface of use case "get entries".
type GetEntriesUseCase interface {
	Run() ([]entity.Entry, error)
}

// GetEntriesUseCase represents the use case of getting entries.
// It implements GetEntriesUseCaseImpl interface.
type getEntriesUseCase struct {
	EntryRepo repository.IEntryRepository
}

// Run exec to the use case.
func (u *getEntriesUseCase) Run() ([]entity.Entry, error) {
	entries, err := u.EntryRepo.GetAll()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get entries entities: %#v", err))
		return nil, err
	}
	return entries, nil
}

// NewGetEntriesUseCase create the use case of getting entries.
func NewGetEntriesUseCase() GetEntriesUseCase {
	return &getEntriesUseCase{EntryRepo: &repository.EntryRepository{}}
}
