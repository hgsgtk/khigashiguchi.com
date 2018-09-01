package datastore_test

import (
	"testing"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/infrastructure/datastore"
	"github.com/google/go-cmp/cmp"
)

func TestEntryStore_GetAll(t *testing.T) {
	tests := []struct {
		name             string
		expectedEntities []entity.Entry
		expectedErr      error
	}{
		{
			name: "get_records",
			expectedEntities: []entity.Entry{
				{
					Title: "ECS(Fargate)で動かすコンテナにSSMからクレデンシャル情報を渡す",
					URL:   "http://khigashigashi.hatenablog.com/entry/2018/08/28/214417",
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := datastore.EntryStore{}
			entries, err := s.GetAll()
			if tt.expectedErr != err {
				t.Errorf("expected err: %#v,\n given: %#v", tt.expectedErr, err)
			}
			if diff := cmp.Diff(tt.expectedEntities, entries); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}
