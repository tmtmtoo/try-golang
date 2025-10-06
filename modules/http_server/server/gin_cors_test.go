package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

func TestCORSMiddleware(t *testing.T) {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	t.Run("OPTIONS", func(t *testing.T) {
		req, _ := http.NewRequest("OPTIONS", "/test", nil)
		req.Header.Set("Origin", "http://localhost")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, 204)
		assert.Equal(t, w.Header().Get("Access-Control-Allow-Origin"), "http://localhost")
		// ...snip...
	})

	t.Run("GET", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Origin", "http://localhost")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, 200)
		assert.Equal(t, w.Header().Get("Access-Control-Allow-Origin"), "http://localhost")
		// ...snip...
	})
}
