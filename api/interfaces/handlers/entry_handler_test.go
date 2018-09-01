package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Khigashiguchi/khigashiguchi.com/api/domain/entity"
	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/handlers"
	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/presenter"
	"github.com/google/go-cmp/cmp"
)

func TestGetEntriesHandler(t *testing.T) {
	tests := []struct {
		name         string
		mUseCaseRun  func() ([]entity.Entry, error)
		expectedCD   int
		expectedBody presenter.GetEntriesResponse
	}{
		{
			name: "respond_any_record",
			mUseCaseRun: func() ([]entity.Entry, error) {
				return []entity.Entry{
					{
						Title: "test title",
						URL:   "http://example.com",
					},
				}, nil
			},
			expectedCD: http.StatusOK,
			expectedBody: presenter.GetEntriesResponse{
				Entities: []entity.Entry{
					{
						Title: "test title",
						URL:   "http://example.com",
					},
				},
			},
		},
		{
			name: "respond_no_record",
			mUseCaseRun: func() ([]entity.Entry, error) {
				return []entity.Entry{}, nil
			},
			expectedCD: http.StatusOK,
			expectedBody: presenter.GetEntriesResponse{
				Entities: []entity.Entry{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/entries", nil)

			m := mUseCase{mRun: tt.mUseCaseRun}
			h := handlers.ExportGetEntriesHandler{UseCase: &m}
			h.Handler(w, r)

			res := w.Result()
			defer res.Body.Close()

			if tt.expectedCD != res.StatusCode {
				t.Errorf("expected: %#v, given #%v", tt.expectedCD, res.StatusCode)
			}
			if expected := "application/json; charset=utf-8"; expected != res.Header.Get("Content-Type") {
				t.Errorf("expected: #%v, given #%v", expected, res.Header.Get("Content-Type"))
			}
			var body presenter.GetEntriesResponse
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Errorf("unexpected error by json.Decode(): %#v", err)
			}
			if diff := cmp.Diff(tt.expectedBody, body); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestGetEntriesHandler_Error(t *testing.T) {
	tests := []struct {
		name         string
		mUseCaseRun  func() ([]entity.Entry, error)
		expectedCD   int
		expectedBody presenter.ApiError
	}{
		{
			name: "respond_any_record",
			mUseCaseRun: func() ([]entity.Entry, error) {
				return []entity.Entry{}, errors.New("test error")
			},
			expectedCD: http.StatusInternalServerError,
			expectedBody: presenter.ApiError{
				Message: "Internal Server Error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/entries", nil)

			m := mUseCase{mRun: tt.mUseCaseRun}
			h := handlers.ExportGetEntriesHandler{UseCase: &m}
			h.Handler(w, r)

			res := w.Result()
			defer res.Body.Close()

			if tt.expectedCD != res.StatusCode {
				t.Errorf("expected: %#v, given #%v", tt.expectedCD, res.StatusCode)
			}
			if expected := "application/json; charset=utf-8"; expected != res.Header.Get("Content-Type") {
				t.Errorf("expected: #%v, given #%v", expected, res.Header.Get("Content-Type"))
			}
			var body presenter.ApiError
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Errorf("unexpected error by json.Decode(): %#v", err)
			}
			if diff := cmp.Diff(tt.expectedBody, body); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

type mUseCase struct {
	mRun func() ([]entity.Entry, error)
}

func (m *mUseCase) Run() ([]entity.Entry, error) {
	return m.mRun()
}
