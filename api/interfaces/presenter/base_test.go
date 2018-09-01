package presenter_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Khigashiguchi/khigashiguchi.com/api/interfaces/presenter"
	"github.com/google/go-cmp/cmp"
)

type testResponse struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func TestRespondJson(t *testing.T) {
	tests := []struct {
		name         string
		body         interface{}
		cd           int
		expectedCD   int
		expectedBody string
	}{
		{
			name: "respond_struct_status_200",
			body: testResponse{
				Field1: "test",
				Field2: 1,
			},
			cd:           http.StatusOK,
			expectedCD:   http.StatusOK,
			expectedBody: "{\"field1\":\"test\",\"field2\":1}\n",
		},
		{
			name:         "respond_empty_status_200",
			body:         struct{}{},
			cd:           http.StatusOK,
			expectedCD:   http.StatusOK,
			expectedBody: "{}\n",
		},
		{
			name:         "respond_status_500",
			body:         struct{}{},
			cd:           http.StatusInternalServerError,
			expectedCD:   http.StatusInternalServerError,
			expectedBody: "{}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			presenter.RespondJson(w, tt.body, tt.cd)

			res := w.Result()
			defer res.Body.Close()

			if tt.expectedCD != res.StatusCode {
				t.Errorf("expected: %#v, given #%v", tt.expectedCD, res.StatusCode)
			}
			if expected := "application/json; charset=utf-8"; expected != res.Header.Get("Content-Type") {
				t.Errorf("expected: #%v, given #%v", expected, res.Header.Get("Content-Type"))
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("unexpected error by ioutil.ReadAll(): %#v", err)
			}
			if diff := cmp.Diff(tt.expectedBody, string(body)); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}

func TestRespondErrorJson(t *testing.T) {
	tests := []struct {
		name         string
		msg          string
		code         int
		expectedBody string
	}{
		{
			name:         "error_status_404",
			msg:          "Not Found",
			code:         http.StatusNotFound,
			expectedBody: "{\"message\":\"Not Found\"}\n",
		},
		{
			name:         "error_status_500",
			msg:          "Internal Server Error",
			code:         http.StatusInternalServerError,
			expectedBody: "{\"message\":\"Internal Server Error\"}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			presenter.RespondErrorJson(w, tt.msg, tt.code)

			res := w.Result()
			defer res.Body.Close()

			if tt.code != res.StatusCode {
				t.Errorf("expected: %#v, given #%v", tt.code, res.StatusCode)
			}
			if expected := "application/json; charset=utf-8"; expected != res.Header.Get("Content-Type") {
				t.Errorf("expected: #%v, given #%v", expected, res.Header.Get("Content-Type"))
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("unexpected error by ioutil.ReadAll(): %#v", err)
			}
			if diff := cmp.Diff(tt.expectedBody, string(body)); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}
