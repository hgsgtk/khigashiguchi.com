package repository

import "github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"

// IEntryRepository is interface of the repository to fetch Entry.
type IEntryRepository interface {
	GetAll(db Executor) ([]entity.Entry, error)
}

// EntryRepository implements IEntryRepository interface.
type EntryRepository struct{}

// GetAll get all entries from data store.
func (s *EntryRepository) GetAll(db Executor) ([]entity.Entry, error) {
	sql := "SELECT title, url FROM entries"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := make([]entity.Entry, 0)
	for rows.Next() {
		entry := entity.Entry{}
		if err := rows.Scan(&entry.Title, &entry.URL); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}
