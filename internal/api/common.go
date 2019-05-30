/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import (
	"encoding/json"
	"errors"
	"math/big"
	"net/http"

	"github.com/eeonevision/avito-pro-test/pkg/rndgen"
)

// Necessary types for calling the methods of random generator.
const (
	TypeGUID     = "guid"
	TypeString   = "string"
	TypeNumber   = "number"
	TypeAlphaNum = "alphanum"
	//TypeRangeValues = "range"
)

type baseResponse struct {
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Result interface{} `json:"result,omitempty"`
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

func getRandomValueByType(t string, length int) (string, error) {
	var res string
	var err error

	switch t {
	case TypeString:
		res, err = rndgen.GetString(length)
		break
	case TypeNumber:
		var tmp *big.Int
		tmp, err = rndgen.GetNumber(length)
		res = tmp.String()
		break
	case TypeGUID:
		res, err = rndgen.GetGUID(length)
		break
	case TypeAlphaNum:
		res, err = rndgen.GetAlphaNum(length)
		break
	default:
		if len(t) == 0 {
			err = errors.New("the type field is empty")
			break
		}
		res, err = rndgen.GetFromValuesRange(t, length) // if type not nil, then it is range type
		break
	}

	return res, err
}
