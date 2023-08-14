package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type PesertaBPJS struct {
	Id               int           `gorm:"type:bigint;primaryKey;autoIncrement;index"`
	NoBPJS           string        `gorm:"type:varchar;primaryKey;index"`
	Name             string        `gorm:"type:varchar;index"`
	BirthDate        time.Time     `gorm:"type:date;default:null"`
	Address          string        `gorm:"type:varchar"`
	FaskesLevel      string        `gorm:"type:varchar"`
	FaskesName       string        `gorm:"type:varchar"`
	OneToManyAnteran []AntreanBPJS `gorm:"foreignKey:NoBPJS"`
}

type PesertaBPJSConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *PesertaBPJSConn) GetPesertaByNoBPJS(noBpjs string) (data *PesertaBPJS, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&PesertaBPJS{NoBPJS: noBpjs}).First(&data)
	if result.Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *PesertaBPJSConn) RegisterPesertaBPJS(dataRegister PesertaBPJS) (err error) {
	if conn.Err != nil {
		return conn.Err
	}

	result := conn.DB.Create(&dataRegister)
	return result.Error
}
