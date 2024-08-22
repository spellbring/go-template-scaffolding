package router

import (
	"fmt"
	"net/http"
	"time"

	"{{bootstrap_template}}/application/adapter/api/health"
	"{{bootstrap_template}}/pkg/log/logger"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type gorillaMux struct {
	router     *mux.Router
	middleware *negroni.Negroni
	log        logger.Logger
	port       Port
	ctxTimeout time.Duration
}

func newGorillaMux(
	log logger.Logger,
	port Port,
	t time.Duration,
) *gorillaMux {
	return &gorillaMux{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		log:        log,
		port:       port,
		ctxTimeout: t,
	}
}

func (g gorillaMux) GetHttpServer() *http.Server {
	g.setAppHandlers(g.router)
	g.middleware.UseHandler(g.router)

	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.middleware,
	}

}

func (g gorillaMux) setAppHandlers(router *mux.Router) {

	api := router
	api.HandleFunc("/health", health.HealthCheck).Methods(http.MethodGet)

	router.Use(SanitizeBodyMiddleware)
}
