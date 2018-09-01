package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// RespondJson return the response.
func RespondJson(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
