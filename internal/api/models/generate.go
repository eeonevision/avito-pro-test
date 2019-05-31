package models

import "errors"

// Generate is implementation of th OK interface to validate fields
// for handlePostGenerate method.
type Generate struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
}

// Validate implements validation rules for Generate struct.
func (gr *Generate) Validate() error {
	if len(gr.Type) == 0 {
		return errors.New("type field is required")
	}
	if gr.Length == 0 {
		return errors.New("length field should be not zero")
	}

	return nil
}
