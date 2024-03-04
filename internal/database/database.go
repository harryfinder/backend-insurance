package database

import (
	"context"
	"github.com/harryfinder/backend-Insurance/internal/models"
)

type Database interface {
	DInsuranceService(ctx context.Context, insuranceData models.VehicleInsurance) error
}
