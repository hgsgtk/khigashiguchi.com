package usecase

import (
	"errors"
	"fmt"
	"os"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/repository"
)

type PostEntriesUseCase interface {
	Run(entry entity.Entry) error
}

type postEntriesUseCase struct {
	DB        repository.Beginner
	EntryRepo repository.IEntryRepository
}

func (u *postEntriesUseCase) Run(entry entity.Entry) error {
	tx, err := u.DB.Begin()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			fmt.Fprintf(os.Stdout, "failed to commit and rollback transaction in postEntriesUseCase.Run(): %s", err)
			return errors.New("internal server error")
		}
		fmt.Fprintf(os.Stdout, "failed to begin transaction in postEntriesUseCase.Run(): %s", err)
		return errors.New("internal server error")
	}

	err = u.EntryRepo.Save(tx, entry)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			fmt.Fprintf(os.Stdout, "failed to commit and rollback transaction in postEntriesUseCase.Run(): %s", err)
			return errors.New("internal server error")
		}
		fmt.Fprintf(os.Stdout, "failed to save entries in postEntriesUseCase.Run(): %s", err)
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			fmt.Fprintf(os.Stdout, "failed to commit and rollback transaction in postEntriesUseCase.Run(): %s", err)
			return errors.New("internal server error")
		}
		fmt.Fprintf(os.Stdout, "failed to commit and rollback transaction in postEntriesUseCase.Run(): %s", err)
		return errors.New("internal server error")
	}

	return nil
}

func NewPostEntriesUseCase(db repository.Beginner) PostEntriesUseCase {
	return &postEntriesUseCase{
		DB:        db,
		EntryRepo: &repository.EntryRepository{}}
}
