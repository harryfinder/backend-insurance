package controller

import (
	"context"
	"github.com/harryfinder/backend-Insurance/internal/models"
)

type Controller interface {
	Serve(context.Context, *models.Configuration) error
	Shutdown(ctx context.Context) error
}
