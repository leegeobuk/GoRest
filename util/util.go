package util

import "net/http"

// CheckErr checks and reply with error message and status code if any
func CheckErr(w http.ResponseWriter, err error, errMsg string, statusCode int) {
	if err != nil {
		http.Error(w, errMsg, statusCode)
	}
}
