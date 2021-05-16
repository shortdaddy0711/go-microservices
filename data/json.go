package data

import (
	"encoding/json"
	"io"
)

// FromJSON deserializes the object from json string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}

// ToJSON serializes the given interface into a json formatted string
func ToJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i)
}
