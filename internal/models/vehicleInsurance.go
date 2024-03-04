package models

type VehicleInsurance struct {
	ID                 string `json:"id"`
	FullName           string `json:"full_name"`
	PolicyID           string `json:"policy_id"`
	Phone              string `json:"phone"`
	VIN                string `json:"vin"`
	InsuranceType      string `json:"insurance_type"`
	PremiumAmount      string `json:"premium_amount"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	Duration           string `json:"duration"`
	RegistrationNumber string `json:"registration_number"`
	EToken             string `json:"e_token"`
}
