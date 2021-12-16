package errors

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Errors struct {
	code    int32
	message string
}

func (e *Errors) Code() int32 {
	return e.code
}

func (e *Errors) Message() string {
	return e.message
}

// MarshalJSON implements the JSON encoding interface
func (e *Errors) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code":    e.code,
		"message": e.message,
	})
}

func (e *Errors) ToError() error {
	return errors.New(fmt.Sprintf("%v: %v", e.Code(), e.Message()))
}

// New creates a new errors instance
func New(code int32, message string, args []any) *Errors {
	if len(args) > 0 {
		return &Errors{code, fmt.Sprintf(message, args...)}
	}
	return &Errors{code, message}
}