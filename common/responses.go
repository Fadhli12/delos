package common

import (
	"encoding/json"
	"net/http"
)

type Responses struct {
	Body  interface{}
	Error error
}

// ResponseJson :
func ResponseJson(w http.ResponseWriter, resp *Responses) {
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != nil {
		request, ok := resp.Error.(*RequestError)
		if ok {
			w.WriteHeader(request.StatusCode)
			json.NewEncoder(w).Encode(request.Err.Error())
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp.Body)
}
