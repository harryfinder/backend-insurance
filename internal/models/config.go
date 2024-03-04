package models

type Configuration struct {
	Version     string `json:"version"`
	Api         string `json:"api"`
	DatabaseDsn string `json:"databaseDsn"`
	MssqlDsn    string `json:"mssqlDsn"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}
