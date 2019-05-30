/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// handlePostGenerate generates random value and it idendtifier.
func (s *server) handlePostGenerate() http.HandlerFunc {
	type request struct {
		Type   string `json:"type"`
		Length int    `json:"length"`
	}

	type response struct {
		ID uint `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			writeJSONResponse(http.StatusBadRequest, err.Error(), nil, w)
			return
		}

		randomVal, err := getRandomValueByType(req.Type, req.Length)
		if err != nil {
			writeJSONResponse(http.StatusBadRequest, err.Error(), nil, w)
			return
		}

		writeJSONResponse(
			http.StatusCreated,
			"The random value was successfully generated",
			response{s.db.Insert(randomVal)},
			w)
	}
}

// handleRetrieve gets the value by id from generate endpoint.
func (s *server) handleGetRetrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		requestID, _ := strconv.Atoi(vars["requestID"])
		value := s.db.GetValueByID(uint(requestID))
		if len(value) == 0 {
			writeJSONResponse(http.StatusNotFound, "The request ID not found", nil, w)
			return
		}

		writeJSONResponse(http.StatusOK, "OK", value, w)
	}
}

// NotFoundHandler returns 404 status code for wrong api routes.
func (s *server) handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSONResponse(http.StatusNotFound, "The resource not found", nil, w)
	}
}
