package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionORM() (*gorm.DB, error) {

	connString := "user=" + os.Getenv("DATABASE_USER") +
		" password=" + os.Getenv("DATABASE_PASSWRD") +
		" host=" + os.Getenv("DATABASE_HOSTNAME") +
		" port=" + os.Getenv("DATABASE_PORT") +
		" dbname=" + os.Getenv("DATABASE_NAME") +
		" sslmode=" + os.Getenv("DATABASE_SSLMODE") +
		" TimeZone=" + os.Getenv("DATABASE_TIMEZONE")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}
