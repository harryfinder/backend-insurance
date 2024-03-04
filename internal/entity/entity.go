package entity

import (
	"context"
	"github.com/harryfinder/backend-Insurance/internal/database"
	"github.com/harryfinder/backend-Insurance/internal/models"
)

type Entity interface {
	EInsurance(ctx context.Context, insurance models.VehicleInsurance) error
}

type entity struct {
	database database.Database
}

func New(database database.Database) Entity {
	return &entity{database: database}

}
