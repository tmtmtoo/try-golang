package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRouter(t *testing.T) {
	tests := []struct {
		method   string
		path     string
		assertFn func(t *testing.T, code int, body string)
	}{
		{
			"GET",
			"/v1/foo",
			func(t *testing.T, code int, body string) {
				assert.Equal(t, code, http.StatusOK)
				assert.Equal(t, body, `[{"id":0},{"id":1}]`)
			},
		},
		{
			"GET",
			"/v1/foo/0",
			func(t *testing.T, code int, body string) {
				assert.Equal(t, code, http.StatusOK)
				assert.Equal(t, body, `{"id":0}`)
			},
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s %s", tt.method, tt.path)
		t.Run(name, func(t *testing.T) {
			router := CreateRouter()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			code := w.Code
			body := w.Body.String()
			tt.assertFn(t, code, body)
		})
	}
}
