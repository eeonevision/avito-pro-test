/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import (
	"encoding/json"
	"net/http"

	"github.com/eeonevision/avito-pro-test/internal/api/models"
)

type baseResponse struct {
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// decode does simple decoding stuff and validation of the struct fields.
func decode(r *http.Request, v models.OK) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate()
}

func writeJSONResponse(code int, message string, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	res, _ := json.Marshal(baseResponse{
		Code:   code,
		Msg:    message,
		Result: data,
	})
	w.Write(res)
}
