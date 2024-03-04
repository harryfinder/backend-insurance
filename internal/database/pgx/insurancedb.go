package pgx

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/harryfinder/backend-Insurance/internal/models"
)

func (d *db) DInsuranceService(ctx context.Context, insuranceData models.VehicleInsurance) error {

	query := `
        INSERT INTO vehicle_insurances (ID, FullName, PolicyID, Phone, VIN, InsuranceType, PremiumAmount, StartDate, EndDate, Duration, RegistrationNumber, EToken) 
        VALUES (@ID, @FullName, @PolicyID, @Phone, @VIN, @InsuranceType, @PremiumAmount, @StartDate, @EndDate, @Duration, @RegistrationNumber, @EToken)
    `

	_, err := d.mssql.Exec(ctx, query,
		sql.Named("ID", insuranceData.ID),
		sql.Named("FullName", insuranceData.FullName),
		sql.Named("PolicyID", insuranceData.PolicyID),
		sql.Named("Phone", insuranceData.Phone),
		sql.Named("VIN", insuranceData.VIN),
		sql.Named("InsuranceType", insuranceData.InsuranceType),
		sql.Named("PremiumAmount", insuranceData.PremiumAmount),
		sql.Named("StartDate", insuranceData.StartDate),
		sql.Named("EndDate", insuranceData.EndDate),
		sql.Named("Duration", insuranceData.Duration),
		sql.Named("RegistrationNumber", insuranceData.RegistrationNumber),
		sql.Named("EToken", insuranceData.EToken))
	if err != nil {
		return fmt.Errorf("ошибка при вставке данных в БД: %w", err)
	}

	return nil
}
