package util

import (
	"encoding/json"
	"io"
	"net/http"
)

// CheckErr checks and reply with error message and status code if any
func CheckErr(w http.ResponseWriter, err error, errMsg string, statusCode int) bool {
	if err != nil {
		http.Error(w, errMsg, statusCode)
		return true
	}
	return false
}

// FromJSON deserializes object in JSON to Go value
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

// ToJSON serializes given interface to JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
