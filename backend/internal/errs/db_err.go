package errs

import (
	"fmt"
)

type DatabaseError struct {
	Message any `json:"msg"`
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("DB error: %d", e.Message)
}

func NewDBError(err error) DatabaseError {
	return DatabaseError{
		Message: err.Error(),
	}
}

func EmptyResult() DatabaseError {
	return DatabaseError{
		Message: "unexpected: no rows in result",
	}
}
