/*
Copyright 2019 Vladislav Dmitriyev.
*/

package endpoint

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PostGenerateHandler generates random value and it idendtifier.
func PostGenerateHandler(w http.ResponseWriter, r *http.Request) {
	var request request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		writeJSONResponse(http.StatusBadRequest, err.Error(), nil, w)
		return
	}

	res, err := getRandomValueByType(request.Type, request.Length)
	if err != nil {
		writeJSONResponse(http.StatusBadRequest, err.Error(), nil, w)
		return
	}

	writeJSONResponse(
		http.StatusCreated,
		"The random value was successfully generated",
		IdempotentDB.Insert(res),
		w)
}

// GetRetrieveHandler gets the value by id from generate endpoint.
func GetRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	requestID, _ := strconv.Atoi(vars["requestID"])
	value := IdempotentDB.GetValueByID(requestID)
	if len(value) == 0 {
		writeJSONResponse(http.StatusNotFound, "The request ID not found", nil, w)
		return
	}

	writeJSONResponse(http.StatusOK, "OK", value, w)
}

// NotFoundHandler returns 404 status code for wrong api routes.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(http.StatusNotFound, "The resource not found", nil, w)
}
