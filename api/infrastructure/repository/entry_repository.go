package repository

import "github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"

// IEntryRepository is interface of the repository to fetch Entry.
type IEntryRepository interface {
	GetAll() ([]entity.Entry, error)
}

// EntryRepository implements IEntryRepository interface.
type EntryRepository struct{}

// GetAll get all entries from data store.
func (s *EntryRepository) GetAll() ([]entity.Entry, error) {
	// FIXME: 仮実装
	return []entity.Entry{
		{
			Title: "ECS(Fargate)で動かすコンテナにSSMからクレデンシャル情報を渡す",
			URL:   "http://khigashigashi.hatenablog.com/entry/2018/08/28/214417",
		},
	}, nil
}
