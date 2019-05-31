package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eeonevision/avito-pro-test/internal/pkg/idempotent"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandlePostGenerate(t *testing.T) {
	// Given.
	type response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	srv := &server{
		db:     idempotent.NewDB(),
		router: mux.NewRouter(),
	}
	srv.routes()

	tests := []struct {
		Name     string
		Type     string
		Length   int
		Expected response
	}{
		{"Empty type field", "", 10, response{Code: http.StatusBadRequest, Msg: "type field is required"}},
		{"Zero length field", "string", 0, response{Code: http.StatusBadRequest, Msg: "length field should be not zero"}},
		{"Generate random string", "string", 10, response{Code: http.StatusCreated, Msg: "the random value was successfully generated"}},
		{"Generate random number", "number", 10, response{Code: http.StatusCreated, Msg: "the random value was successfully generated"}},
		{"Generate random GUID", "guid", 10, response{Code: http.StatusCreated, Msg: "the random value was successfully generated"}},
		{"Generate random alphanum", "alphanum", 10, response{Code: http.StatusCreated, Msg: "the random value was successfully generated"}},
		{"Generate random from range", "abcef", 10, response{Code: http.StatusCreated, Msg: "the random value was successfully generated"}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			// Given request.
			jsonBody := []byte(fmt.Sprintf(`{"type":"%s","length":%d}`, tc.Type, tc.Length))

			// When.
			req, err := http.NewRequest(http.MethodPost, "/api/v1/generate", bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, req)

			var res response
			decoder := json.NewDecoder(w.Body)
			assert.NoError(t, decoder.Decode(&res))

			// Then.
			assert.Equal(t, tc.Expected.Code, w.Code)                                           // Check status code.
			assert.Equal(t, "application/json; charset=UTF-8", w.HeaderMap.Get("Content-Type")) // Check header content type.

			// Check json payload.
			assert.Equal(t, tc.Expected.Code, res.Code)
			assert.Equal(t, tc.Expected.Msg, res.Msg)
		})
	}
}

func TestHandleRetreive(t *testing.T) {
	// Given.
	type result struct {
		Value string `json:"value"`
	}
	type response struct {
		Code   int     `json:"code"`
		Msg    string  `json:"msg"`
		Result *result `json:"result"`
	}

	testValue := "test value"

	db := idempotent.NewDB()
	id := db.Insert(testValue)

	srv := &server{
		db:     db,
		router: mux.NewRouter(),
	}
	srv.routes()

	// When.
	tests := []struct {
		Name     string
		ID       uint
		Expected response
	}{
		{"Pass not existing ID", id + 1, response{Code: http.StatusNotFound, Msg: "the requested ID not found", Result: nil}},
		{"Pass existing ID", id, response{Code: http.StatusOK, Msg: "OK", Result: &result{Value: testValue}}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/retrieve/%d", tc.ID), nil)
			assert.NoError(t, err)
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, req)

			var res response
			decoder := json.NewDecoder(w.Body)
			assert.NoError(t, decoder.Decode(&res))

			// Then.
			assert.Equal(t, tc.Expected.Code, w.Code)                                           // Check status code.
			assert.Equal(t, "application/json; charset=UTF-8", w.HeaderMap.Get("Content-Type")) // Check header content type.

			// Check json payload.
			assert.Equal(t, tc.Expected.Code, res.Code)
			assert.Equal(t, tc.Expected.Msg, res.Msg)
			assert.Equal(t, tc.Expected.Result, res.Result)
		})
	}
}
