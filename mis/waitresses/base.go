package waitresses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse map[string]interface{}

func (resp JsonResponse) String() (s string) {
	bytes, err := json.Marshal(resp)
	if err != nil {
		s = ""
		return
	}
	s = string(bytes)
	return
}

func (resp JsonResponse) Output(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, resp)
}
