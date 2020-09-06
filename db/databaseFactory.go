package db

import (
	"os"

	"github.com/fajryhamzah/mkpay-user/db/drivers/psql"
)

func get(driver string, connectionString string) DbInterface {
	var db DbInterface

	switch driver {
	case PSQL:
		db = psql.Driver{Connection: connectionString}
	default:
		panic("Driver not supported")
	}

	return db
}

//GetInstance instance of real database
func GetInstance() DbInterface {
	return get(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONNECTION_STRING"))
}

//GetTestInstance instance of testing database
func GetTestInstance() DbInterface {
	return get(os.Getenv("DB_DRIVER"), os.Getenv("DB_TEST_CONNECTION_STRING"))
}
