package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AntreanBPJS struct {
	Id                  int         `gorm:"type:bigint;primaryKey;autoIncrement;index"`
	NoBPJS              string      `gorm:"type:varchar;primaryKey;index"`
	OneToOnePesertaBPJS PesertaBPJS `gorm:"foreignKey:NoBPJS"`
	DateRegister        time.Time   `gorm:"type:date;default:null;default:CURRENT_DATE;index"`
	TimeRegister        time.Time   `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
	PoliId              int         `gorm:"type:integer;index"`
	OneToOnePoli        Poli        `gorm:"foreignKey:Id"`
	NoAntrean           int         `gorm:"type:integer;index"`
	TypePasien          int         `gorm:"type:integer;index"`
	OneToOneTipePasien  TypePasien  `gorm:"foreignKey:Id"`
	Status              bool        `gorm:"type:boolean"`
}

type AntreanBPJSConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *PesertaBPJSConn) AntreanGetPesertaByNoBPJS(noBpjs string) (data *PesertaBPJS, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&PesertaBPJS{NoBPJS: noBpjs}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *AntreanBPJSConn) InsertPendaftaran(dataRegister AntreanBPJS) (err error) {
	if conn.Err != nil {
		return conn.Err
	}

	result := conn.DB.Create(&dataRegister)
	return result.Error
}

func (conn *AntreanBPJSConn) CallPendaftaran() (noAntrean string, err error) {
	if conn.Err != nil {
		return "", conn.Err
	}
	var data AntreanBPJS
	result := conn.DB.Model(&AntreanBPJS{DateRegister: time.Now(), Status: true}).Order("Id ASC").First(&data)
	if result.Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", result.Error
	}
	noAntrean = fmt.Sprintf("%s%04d", data.OneToOnePoli.AntreanCode, data.NoAntrean)
	return noAntrean, nil
}

func (conn *AntreanBPJSConn) DeactiveStatusAntrean(antreanId int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}

	result := conn.DB.Model(&AntreanBPJS{}).Where(AntreanBPJS{Id: antreanId}).Update("Status", true)
	return result.Error
}

func (conn *AntreanBPJSConn) GetLastNumberAntrean(poliId int) (antrean int, err error) {
	if conn.Err != nil {
		return 0, conn.Err
	}

	var data AntreanBPJS
	result := conn.DB.Model(&AntreanBPJS{PoliId: poliId}).Order("NoAntrean desc").First(&data)
	if result.Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, nil
		}
		return 0, result.Error
	}
	return data.NoAntrean + 1, nil
}
