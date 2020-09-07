package api

import (
	"net/http"

	"github.com/koifish082/golang-api-layered/src/app/interfaces/api/middleware"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/helper"

	v1 "github.com/koifish082/golang-api-layered/src/app/interfaces/api/v1"

	"github.com/go-chi/chi"
	"github.com/koifish082/golang-api-layered/src/app/library/log"
)

// App struct
type App struct {
}

// NewApp factory returns an instantiated App
func NewApp() *App {
	return &App{}
}

// Router provides api router
func (a *App) Router() http.Handler {
	router := chi.NewRouter()
	m := middleware.NewMiddleware()
	router.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("hello world!")
		helper.Succeed(w, "welcome!")
	})

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/search", v1.NewSearchResources(m).Routes())
	})

	return router
}
