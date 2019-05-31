/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import "net/http"

func (s *server) routes() {
	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound())

	api := s.router.PathPrefix("/api").Subrouter()

	v1 := api.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/generate", s.handlePostGenerate()).Methods(http.MethodPost)
	v1.HandleFunc("/retrieve/{requestID}", s.handleGetRetrieve()).Methods(http.MethodGet)
}
