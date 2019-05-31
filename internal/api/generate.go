/*
Copyright 2019 Vladislav Dmitriyev.
*/

package api

import (
	"math/big"
	"net/http"

	"github.com/eeonevision/avito-pro-test/internal/api/models"
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
		res, err = rndgen.GetFromValuesRange(t, length) // if type not nil, then it is range type
		break
	}

	return res, err
}

// handlePostGenerate generates random value and it idendtifier.
func (s *server) handlePostGenerate() http.HandlerFunc {
	type response struct {
		ID uint `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req models.Generate
		if err := decode(r, &req); err != nil {
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
			"the random value was successfully generated",
			response{s.db.Insert(randomVal)},
			w)
	}
}
