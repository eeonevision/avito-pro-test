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
	// Given
	type response struct {
		Code   int    `json:"code"`
		Msg    string `json:"msg"`
		Result struct {
			ID int `json:"id"`
		} `json:"result"`
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
		{"Generate random string", "string", 10, response{Code: 201, Msg: "The random value was successfully generated"}},
		{"Generate random number", "number", 10, response{Code: 201, Msg: "The random value was successfully generated"}},
		{"Generate random GUID", "GUID", 10, response{Code: 201, Msg: "The random value was successfully generated"}},
		{"Generate random alphanum", "alphanum", 10, response{Code: 201, Msg: "The random value was successfully generated"}},
		{"Generate random from range", "range[abcef]", 10, response{Code: 201, Msg: "The random value was successfully generated"}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			// Given request
			jsonBody := []byte(fmt.Sprintf(`{"type":"%s","length":%d}`, tc.Type, tc.Length))

			// When
			req, err := http.NewRequest(http.MethodPost, "/api/v1/generate", bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, req)

			var res response
			decoder := json.NewDecoder(w.Body)
			assert.NoError(t, decoder.Decode(&res))

			// Then
			assert.Equal(t, http.StatusCreated, w.Code)                                         // check status code
			assert.Equal(t, "application/json; charset=UTF-8", w.HeaderMap.Get("Content-Type")) // check header content type

			// Check json payload
			assert.Equal(t, tc.Expected.Code, res.Code)
			assert.Equal(t, tc.Expected.Msg, res.Msg)
			assert.NotZero(t, res.Result.ID)
		})
	}
}
