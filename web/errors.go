package web

import (
	"net/http"
	"runtime"
	"strings"
)

// an application error
type apiError struct {
	code    int
	errType string
	message string
	orig    error
}

func (a apiError) Error() string {
	return a.message
}

func buildErr(code int, orig error, message string) error {
	errType := "UNKNOWNTYPE"
	if pc, _, _, ok := runtime.Caller(1); ok {
		if f := runtime.FuncForPC(pc); f != nil {
			parts := strings.Split(f.Name(), ".")
			errType = parts[len(parts)-1]
		}
	}
	return &apiError{
		code:    code,
		orig:    orig,
		message: message,
		errType: errType,
	}
}

func errDecodingJSON(err error) error {
	return buildErr(http.StatusBadRequest, err, "Invalid json")
}
