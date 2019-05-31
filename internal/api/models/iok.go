package models

// OK interface represents the types
// capable to validate themselves.
type OK interface {
	Validate() error
}
