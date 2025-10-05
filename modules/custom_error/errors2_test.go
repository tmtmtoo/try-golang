package customerror

import (
	"errors"
	"testing"
)

func TestErrorAs2(t *testing.T) {
	err := NewValidationError2("not found", errors.New("this is error"))

	var ce *CustomError
	if errors.As(err, &ce) {
		switch ce.Type {
		case ValidationErrorType:
			t.Logf("msg: %s", ce.msg)
		default:
			t.Error("fucking unknown error")
		}
	} else {
		t.Errorf("fuck")
	}
}
