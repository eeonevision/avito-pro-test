/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import (
	"net/http"
)

// NotFoundHandler returns 404 status code for wrong api routes.
func (s *server) handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSONResponse(http.StatusNotFound, "the resource not found", nil, w)
	}
}
