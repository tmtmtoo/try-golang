package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Foo struct {
	ID uint64 `json:"id"`
}

var foos = []Foo{
	{ID: 0},
	{ID: 1},
}

func CreateRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/foo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, foos)
	})
	v1.GET("/foo/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		var foo *Foo
		for i := range foos {
			if foos[i].ID == id {
				foo = &foos[i]
				break
			}
		}

		if foo == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		} else {
			ctx.JSON(http.StatusOK, foo)
		}
	})

	return router
}
