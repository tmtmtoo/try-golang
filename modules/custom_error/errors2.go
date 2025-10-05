package customerror

import "fmt"

type ErrorType int

const (
	ValidationErrorType ErrorType = iota
	ResourceErrorType
	ConflictErrorType
	ExternalErrorType
	SystemErrorType
)

func (t ErrorType) String() string {
	switch t {
	case ValidationErrorType:
		return "Validation"
	case ResourceErrorType:
		return "Resource"
	case ConflictErrorType:
		return "Conflict"
	case ExternalErrorType:
		return "External"
	case SystemErrorType:
		return "System"
	default:
		return "Unknown"
	}
}

type CustomError struct {
	Type ErrorType
	msg  string
	err  error
}

func (e *CustomError) Error() string {
	if e.err == nil {
		return fmt.Sprintf("[%s] %s", e.Type, e.msg)
	}
	return fmt.Sprintf("[%s] %s: %v", e.Type, e.msg, e.err)
}

func (e *CustomError) Unwrap() error {
	return e.err
}

func NewValidationError2(msg string, err error) *CustomError {
	return &CustomError{Type: ValidationErrorType, msg: msg, err: err}
}

func NewResourceError2(msg string, err error) *CustomError {
	return &CustomError{Type: ResourceErrorType, msg: msg, err: err}
}

func NewConflictError2(msg string, err error) *CustomError {
	return &CustomError{Type: ConflictErrorType, msg: msg, err: err}
}

func NewExternalError2(msg string, err error) *CustomError {
	return &CustomError{Type: ExternalErrorType, msg: msg, err: err}
}

func NewSystemError2(msg string, err error) *CustomError {
	return &CustomError{Type: SystemErrorType, msg: msg, err: err}
}
