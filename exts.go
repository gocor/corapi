package corapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// DecodeBody from json into a model
func DecodeBody(ctx context.Context, req *http.Request, m interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(m); err != nil {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	return nil
}

// WriteError ...
func WriteError(ctx context.Context, statusCode int, res http.ResponseWriter) {
	m := CoreResponse{
		ErrorCode:    statusCode,
		ErrorMessage: http.StatusText(statusCode),
	}
	write(ctx, res, m, statusCode)
}

// Write ...
func Write(ctx context.Context, res http.ResponseWriter, m interface{}) {
	write(ctx, res, m, http.StatusOK)
}

// WriteWithStatus ...
func WriteWithStatus(ctx context.Context, res http.ResponseWriter, m interface{}, statusCode int) {
	write(ctx, res, m, statusCode)
}

func write(ctx context.Context, res http.ResponseWriter, m interface{}, statusCode int) {
	if statusCode < 200 || statusCode >= 300 {
		res.WriteHeader(statusCode)
	}

	if err := json.NewEncoder(res).Encode(m); err != nil {
		panicWrite(ctx, res, errors.WithStack(err))
	}
}

func panicWrite(ctx context.Context, res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(http.StatusText(http.StatusInternalServerError)))

	// panic, hoping recoverery will log this
	panic(err)
}
