/*
Copyright 2019 Vladislav Dmitriyev.
*/

// Package rndgen provides utilities to generate
// random values of different types with specified length.
package rndgen

import (
	"crypto/rand"
	"io"
	"math/big"

	"github.com/google/uuid"
)

// random function
// TODO: Check if it's right or notm due the convention in Go codestyle.
var rander = rand.Reader

// const values of alphabets for random generator.
const (
	alphabetBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes   = "0123456789"
	alphanumBytes = numberBytes + alphabetBytes
)

// GetString returns random generated string by specified length in argument.
func GetString(length int) (string, error) {
	if length == 0 {
		return "", buildError("GetString", "provided zero length to generate value")
	}

	res, err := generate(alphabetBytes, length)

	return string(res), err
}

// GetNumber generates random integer value with specified length in argument.
func GetNumber(length int) (*big.Int, error) {
	var zero *big.Int

	if length == 0 {
		return zero, buildError("GetNumber", "provided zero length to generate value")
	}

	res, err := generate(numberBytes, length)
	if err != nil {
		return zero, err
	}

	num := big.NewInt(0)
	num.SetString(string(res), 10)

	return num, err
}

// GetGUID generates random UUIDv4 value with specified length in argument.
// NOTE: if the length is greather than max UUID lenght, then the full UUID will be returned.
func GetGUID(length int) (string, error) {
	uuid.SetRand(rander)

	u, err := uuid.NewRandom()
	if err != nil {
		return "", buildError("GetGUID", err.Error())
	}

	result := u.String()
	if length < len(result) && length > 0 {
		result = result[:length]
	}

	return result, nil
}

// GetAlphaNum returns alphanumerical random value with specified length.
func GetAlphaNum(length int) (string, error) {
	if length == 0 {
		return "", buildError("GetAlphaNum", "provided zero length to generate value")
	}

	res, err := generate(alphanumBytes, length)

	return string(res), err
}

// GetFromValuesRange returns alphanumerical random value with specified length
// using provided alphabet in 'values' argument of the method.
func GetFromValuesRange(values string, length int) (string, error) {
	if len(values) == 0 {
		return "", buildError("GetFromValuesRange", "provided zero alphabet values to generate value")
	}

	if length == 0 {
		return "", buildError("GetFromValuesRange", "provided zero length to generate value")
	}

	res, err := generate(values, length)

	return string(res), err
}

// SetRand sets the random number generator to r, which implements io.Reader.
func SetRand(r io.Reader) {
	if r == nil {
		rander = rand.Reader
		return
	}
	rander = r
}
