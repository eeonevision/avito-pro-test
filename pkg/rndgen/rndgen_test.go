/*
Copyright 2019 Vladislav Dmitriyev.
*/

package rndgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type randomMock struct {
	mock *mock.Mock
	seed []byte
}

func (rm randomMock) Read(p []byte) (n int, err error) {
	return copy(p, rm.seed), nil
}

// getMockRnd return mocked random number generator
func getMockRnd(seed []byte) randomMock {
	rnd := randomMock{&mock.Mock{}, seed}
	rnd.mock.On("Read", rnd.seed).Return(0, nil)

	return rnd
}

func TestGetString(t *testing.T) {
	t.Parallel()

	SetRand(getMockRnd([]byte("100")))

	t.Run("Output is equal", func(ts *testing.T) {
		res, err := GetString(6)
		assert.NoError(ts, err)
		assert.Equal(ts, "XWWaaa", res)
	})

	t.Run("Zero length", func(ts *testing.T) {
		_, err := GetString(0)
		assert.Equal(ts, buildError("GetString", "provided zero length to generate value").Error(), err.Error())
	})
}

func TestGetNumber(t *testing.T) {
	SetRand(getMockRnd([]byte("100")))

	// Table Driven tests
	tests := []struct {
		name     string
		length   int
		expected int
		errStr   string
	}{
		{"Output is equal", 3, 100, ""},
		{"Zero length", 0, 0, buildError("GetNumber", "provided zero length to generate value").Error()},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := GetNumber(tc.length)
			if err != nil {
				assert.Equal(t, err.Error(), tc.errStr)
			}
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestGetGUID(t *testing.T) {
	SetRand(getMockRnd([]byte("100")))

	// Table Driven tests
	tests := []struct {
		name     string
		length   int
		expected string
	}{
		{"Output is equal", 0, "31303031-3030-4130-b031-303031303031"},
		{"Short output is equal", 5, "31303"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := GetGUID(tc.length)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestGetAlphaNum(t *testing.T) {
	t.Parallel()

	SetRand(getMockRnd([]byte("100")))

	t.Run("Output is equal", func(t *testing.T) {
		res, err := GetAlphaNum(6)
		assert.NoError(t, err)
		assert.Equal(t, "NMM000", res)
	})

	t.Run("Zero length", func(t *testing.T) {
		_, err := GetAlphaNum(0)
		assert.Equal(t, buildError("GetAlphaNum", "provided zero length to generate value").Error(), err.Error())
	})
}

func TestGetFromValuesRange(t *testing.T) {
	SetRand(getMockRnd([]byte("100")))

	// Table Driven tests
	tests := []struct {
		name     string
		alphabet string
		length   int
		expected string
		errStr   string
	}{
		{"Output is equal", "01abc", 5, "10000", ""},
		{"Zero alphabet values", "", 5, "", buildError("GetFromValuesRange", "provided zero alphabet values to generate value").Error()},
		{"Zero length", "01abc", 0, "", buildError("GetFromValuesRange", "provided zero length to generate value").Error()},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := GetFromValuesRange(tc.alphabet, tc.length)
			if err != nil {
				assert.Equal(t, err.Error(), tc.errStr)
			}
			assert.Equal(t, tc.expected, res)
		})
	}
}
