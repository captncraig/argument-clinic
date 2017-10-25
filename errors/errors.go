package errors

import (
	"net/http"
	"runtime"
	"strings"
)

// an application error
type ApiError struct {
	Code      int    // http response code
	ErrType   string // some identifier for this class
	Message   string
	Orig      error // original error
	CallerPCs []uintptr
}

func (a ApiError) Error() string {
	return a.Message
}

func (a ApiError) Callers() []uintptr {
	return a.CallerPCs
}

func buildErr(code int, orig error, message string) error {
	// use calling function name for type identifier
	errType := "UNKNOWNTYPE"
	if pc, _, _, ok := runtime.Caller(1); ok {
		if f := runtime.FuncForPC(pc); f != nil {
			parts := strings.Split(f.Name(), ".")
			errType = parts[len(parts)-1]
		}
	}
	pcs := make([]uintptr, 20)
	n := runtime.Callers(3, pcs)

	return &ApiError{
		Code:      code,
		Orig:      orig,
		Message:   message,
		ErrType:   errType,
		CallerPCs: pcs[:n],
	}
}

func ErrDecodingJSON(err error) error {
	return buildErr(http.StatusBadRequest, err, "Invalid json")
}
