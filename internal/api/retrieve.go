/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// handleRetrieve gets the value by id from generate endpoint.
func (s *server) handleGetRetrieve() http.HandlerFunc {
	type response struct {
		Value string `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		requestID, _ := strconv.Atoi(vars["requestID"])
		value := s.db.GetValueByID(uint(requestID))
		if len(value) == 0 {
			writeJSONResponse(http.StatusNotFound, "the requested ID not found", nil, w)
			return
		}

		writeJSONResponse(http.StatusOK, "OK", response{value}, w)
	}
}
