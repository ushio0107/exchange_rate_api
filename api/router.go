package api

import (
	"github.com/gin-gonic/gin"
)

func (a *API) NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/convert", a.handler)

	return r
}
