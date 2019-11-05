package errors

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestErrorHandler(t *testing.T) {
	type args struct {
		response *echo.Response
		message  string
		status   int
		err      error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrorHandler(tt.args.response, tt.args.message, tt.args.status, tt.args.err)
		})
	}
}
