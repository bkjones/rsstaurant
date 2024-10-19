package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an api key from the headers of an http request
// example:
// Authorization: ApiKey <insert api key here>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid auth info format")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
