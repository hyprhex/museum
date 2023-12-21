package api

import (
	"encoding/json"
	"net/http"

	"github.com/hyprhex/museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exhinition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhinition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exhinition)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Ok"))
	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}
