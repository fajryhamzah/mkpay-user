package db

import (
	"os"

	"github.com/fajryhamzah/mkpay-user/db/drivers/psql"
)

func get(driver string, connectionString string) DbInterface {
	var db DbInterface

	switch driver {
	case PSQL:
		db = psql.Driver{connectionString}
	default:
		panic("Driver not supported")
	}

	return db
}

func GetInstance() DbInterface {
	return get(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONNECTION_STRING"))
}
