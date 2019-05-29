/*
Copyright 2019 Vladislav Dmitriyev.
*/

package endpoint

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eeonevision/avito-pro-test/internal/pkg/idempotent"
	"github.com/eeonevision/avito-pro-test/pkg/rndgen"
)

// Necessary types for calling the methods of random generator.
const (
	TypeGUID        = "guid"
	TypeString      = "string"
	TypeNumber      = "number"
	TypeAlphaNum    = "alphanum"
	TypeRangeValues = "range"
)

// IdempotentDB represents simple database for keeping generated id-result pairs.
// There is not right solution, but for start is ok.
// TODO: convert it to handler/endpoint struct with field DB in it.
var IdempotentDB *idempotent.DB

type requestID struct {
	ID int `json:"id"`
}

type request struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
}

type response struct {
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func writeJSONResponse(code int, message string, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	trs, _ := json.Marshal(response{
		Code:   code,
		Msg:    message,
		Result: data,
	})
	w.Write(trs)
}

func getRandomValueByType(t string, length int) (string, error) {
	var res string
	var err error

	switch t {
	case TypeString:
		res, err = rndgen.GetString(length)
	case TypeNumber:
		var tmp int
		tmp, err = rndgen.GetNumber(length)
		res = strconv.Itoa(tmp)
	case TypeGUID:
		res, err = rndgen.GetGUID(length)
	case TypeAlphaNum:
		res, err = rndgen.GetAlphaNum(length)
	}

	return res, err
}
