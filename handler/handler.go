package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rudbast/tax-calc/api"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Data  interface{} `json:"data"`
		Error string      `json:"error,omitempty"`
	}{}

	data, err := fn(w, r)
	if err != nil {
		log.Println("Process request error:", err.Err)
		resp.Error = err.Message

		// Modify status code on error, default 200 (OK).
		w.WriteHeader(err.StatusCode)
	} else {
		resp.Data = data
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println("Encode response error:", err)
	}
}
