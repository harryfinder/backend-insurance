package usecase

import (
	"context"
	"errors"
	"github.com/harryfinder/backend-Insurance/internal/models"
)

func (u *usecase) UInsurance(ctx context.Context, insurance models.VehicleInsurance) error {

	if err := validate(insurance); err != nil {
		return errors.New("Ошибка во время валидации ")
	}

	if err := u.entity.EInsurance(ctx, insurance); err != nil {
		return errors.New("ошибка во после валидации и передачи данных для создание токена")
	}

	return nil

}
func validate(insurance models.VehicleInsurance) error {
	if insurance.FullName == "" {
		return errors.New("ФИО не может быть пустым")
	}
	if insurance.Phone == "" {
		return errors.New("Номер телефона не может быть пустым")
	}
	if insurance.RegistrationNumber == "" {
		return errors.New("Регистрационный номер машины не может быть пустым")
	}
	if insurance.VIN == "" {
		return errors.New("VIN не может быть пустым")
	}
	if insurance.PremiumAmount == "" {
		return errors.New("Цена не может быть пустым")
	}

	return nil
}
