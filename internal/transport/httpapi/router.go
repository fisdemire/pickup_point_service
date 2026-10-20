package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Deps struct {
	Auth   AuthService
	Tokens TokenParser
}

func NewRouter(d Deps) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, errorDTO{Message: "not found"}) })

	h := &Handler{auth: d.Auth}
	r.POST("/dummyLogin", h.dummyLogin)
	r.POST("/register", h.register)
	r.POST("/login", h.login)
	return r
}
