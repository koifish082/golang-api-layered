package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koifish082/golang-api-layered/src/app/library/log"
)

type App struct {
}

// NewApp factory returns an instantiated App
func NewApp() *App {
	return &App{}
}

func (a *App) Router() *gin.Engine {
	router := gin.New()
	router.GET("/welcome", func(c *gin.Context) {
		log.Infoln("hello world!")
		c.String(http.StatusOK, "hello world")
	})
	gin.Default()
	return router
}
