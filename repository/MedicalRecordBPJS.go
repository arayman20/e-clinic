package repository

import (
	"gorm.io/gorm"
)

type MedicalRecordBPJSConn struct {
	DB  *gorm.DB
	Err error
}
