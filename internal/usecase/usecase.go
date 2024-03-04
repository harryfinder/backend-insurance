package usecase

import (
	"context"
	"github.com/harryfinder/backend-Insurance/internal/entity"
	"github.com/harryfinder/backend-Insurance/internal/models"
)

type UseCase interface {
	UInsurance(ctx context.Context, insurance models.VehicleInsurance) error
}

type usecase struct {
	entity entity.Entity
}

func New(entity entity.Entity) UseCase {
	return &usecase{entity: entity}

}
