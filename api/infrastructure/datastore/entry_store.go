package datastore

import "github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"

// EntryStore implements EntryRepository interface.
type EntryStore struct{}

// GetAll get all entries from data store.
func (s *EntryStore) GetAll() ([]entity.Entry, error) {
	// FIXME: 仮実装
	return []entity.Entry{
		{
			Title: "ECS(Fargate)で動かすコンテナにSSMからクレデンシャル情報を渡す",
			URL:   "http://khigashigashi.hatenablog.com/entry/2018/08/28/214417",
		},
	}, nil
}
