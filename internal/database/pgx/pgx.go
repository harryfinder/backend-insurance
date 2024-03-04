package pgx

import (
	"github.com/harryfinder/backend-Insurance/internal/database"
	"github.com/harryfinder/backend-Insurance/pkg/storage"
)

type db struct {
	mssql storage.Database
}

func New(msSql storage.Database) database.Database {
	return &db{mssql: msSql}
}
