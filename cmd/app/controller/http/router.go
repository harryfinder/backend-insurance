package http

import (
	"context"
	"github.com/harryfinder/backend-Insurance/cmd/app/controller"
	"github.com/harryfinder/backend-Insurance/internal/models"
	"github.com/harryfinder/backend-Insurance/internal/usecase"
	pkghttp "github.com/harryfinder/backend-Insurance/pkg/controller/http"

	"net/http"
)

type server struct {
	usecase usecase.UseCase
	srv     pkghttp.Server
}

func NewController(usecase usecase.UseCase, srv pkghttp.Server) controller.Controller {
	return &server{usecase: usecase, srv: srv}
}

func (s *server) Serve(ctx context.Context, config *models.Configuration) error {
	return s.srv.Serve(ctx, config, []pkghttp.Route{
		//PING
		{Method: http.MethodGet, Path: "/ping", Handler: s.ping},
		{Method: http.MethodPost, Path: "/insurance", Handler: s.Insurance},
	})
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
