package api

import (
	"encoding/base64"
	"strings"
)

// NOTE: these are only used in cleartext format in development scenarios.  At compile-time a
// known token is substituted for the password.
const (
	AuthUsername = "klearwave"
	AuthPassword = "developer"
)

type Authorization struct {
	Authorization string `header:"Authorization"`
}

// Authorized parses the Authorization header for the 'Basic' prefix and a
// base64-encoded string.  It returns an error if the header is missing,
// or the base64-encoded string is not valid.  The base64-encoded string
// should match hard-coded values for the username and password and is set
// at compile time (for production) and as a constant above (for development).
func (auth *Authorization) Authorized() (bool, error) {
	if auth.Authorization == "" {
		return false, nil
	}

	if !strings.HasPrefix(auth.Authorization, "Basic ") {
		return false, nil
	}

	// decode the base64-encoded string
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth.Authorization, "Basic "))
	if err != nil {
		return false, err
	}

	return strings.TrimSpace(string(decoded)) == AuthUsername+":"+AuthPassword, nil
}
