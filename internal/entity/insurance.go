package entity

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/harryfinder/backend-Insurance/internal/models"
	"time"
)

func (e *entity) EInsurance(ctx context.Context, insurance models.VehicleInsurance) error {

	sourceString := insurance.RegistrationNumber + time.Now().String()
	hasher := sha256.New()
	hasher.Write([]byte(sourceString))
	token := hex.EncodeToString(hasher.Sum(nil))

	insurance.EToken = token

	return e.database.DInsuranceService(ctx, insurance)
}
