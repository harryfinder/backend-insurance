package http

import (
	"context"
	"fmt"
	"github.com/harryfinder/backend-Insurance/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type server struct {
	srv *http.Server
}

type Server interface {
	Serve(context.Context, *models.Configuration, []Route) error
	Shutdown(ctx context.Context) error
}

func NewServer() Server {
	return &server{
		srv: &http.Server{
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 120 * time.Second,
		},
	}
}

type Route struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (s *server) Serve(ctx context.Context, cfg *models.Configuration, routes []Route) error {
	//localhost:8080
	s.srv.Addr = cfg.Host + ":" + cfg.Port
	fmt.Println(s.srv.Addr)
	handler := httprouter.New()

	for _, route := range routes {
		handler.Handle(route.Method, route.Path, route.Handler)
	}
	s.srv.Handler = handler
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
