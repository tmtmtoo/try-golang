package customerror

type (
	ValidationError struct {
		msg string
		err error
	}
	ResourceError struct {
		msg string
		err error
	}
	ConflictError struct {
		msg string
		err error
	}
	ExternalError struct {
		msg string
		err error
	}
	SystemError struct {
		msg string
		err error
	}
)

func NewValidationError(msg string, err error) *ValidationError {
	return &ValidationError{msg, err}
}

func (e *ValidationError) Error() string {
	return e.msg
}

func (e *ValidationError) Unwrap() error {
	return e.err
}

func NewResourceError(msg string, err error) *ResourceError {
	return &ResourceError{msg, err}
}

func (e *ResourceError) Error() string {
	return e.msg
}

func (e *ResourceError) Unwrap() error {
	return e.err
}

func NewConflictError(msg string, err error) *ConflictError {
	return &ConflictError{msg, err}
}

func (e *ConflictError) Error() string {
	return e.msg
}

func (e *ConflictError) Unwrap() error {
	return e.err
}

func NewExternalError(msg string, err error) *ExternalError {
	return &ExternalError{msg, err}
}

func (e *ExternalError) Error() string {
	return e.msg
}

func (e *ExternalError) Unwrap() error {
	return e.err
}

func NewSystemError(msg string, err error) *SystemError {
	return &SystemError{msg, err}
}

func (e *SystemError) Error() string {
	return e.msg
}

func (e *SystemError) Unwrap() error {
	return e.err
}
