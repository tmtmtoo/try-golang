package customerror

import (
	"errors"
	"testing"
)

func TestErrorAs(t *testing.T) {
	err := NewValidationError("not found", errors.New("this is error"))

	var validationError *ValidationError
	if !errors.As(err, &validationError) {
		t.Errorf("fuck")
	}

	if validationError.msg != "not found" {
		t.Errorf("fucking message")
	}
}
