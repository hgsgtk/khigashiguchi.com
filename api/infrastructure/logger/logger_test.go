package logger

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/google/go-cmp/cmp"
)

func TestLogger_Error(t *testing.T) {
	tests := []struct{
		name string
		inputMsg string
		expectedMsg string
	}{
		{
			name: "log_msg",
			inputMsg: "test error message",
			expectedMsg: "{\"level\":\"error\",\"msg\":\"test error message\"}\n",
		},
	}
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("unexpected by os.Pipe(): %#v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer = w
			Error(tt.inputMsg)
			w.Close()

			actual, err := ioutil.ReadAll(r)
			if err != nil {
				t.Errorf("unexpected by ioutil.ReadAll(): %#v", err)
			}

			if diff := cmp.Diff(tt.expectedMsg, string(actual)); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
		})
	}
}