package repository

import (
	"errors"
	"klinik/basemodel"

	"gorm.io/gorm"
)

type TypePasien struct {
	Id       int    `gorm:"type:bigint;primaryKey;autoIncrement;index"`
	CodeName string `gorm:"varchar"`
	Name     string `gorm:"varchar"`
	Status   bool   `gorm:"type:boolean;default:True"`
}

type TypePasienConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *TypePasienConn) DataTypePasien(pagination basemodel.TypePasienPaginationBaseModel) (data []basemodel.TypePasienDataBaseModel, count int64, err error) {
	if conn.Err != nil {
		return nil, 0, conn.Err
	}

	db := conn.DB.Model(&TypePasien{})
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	result := db.
		Order("Id DESC").
		Limit(pagination.Size).
		Offset((pagination.Page - 1) * pagination.Size).
		Find(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, result.Error
	}
	return data, count, nil

}

func (conn *TypePasienConn) GetTypePasienByCodeName(codeName string) (data *basemodel.TypePasienDataBaseModel, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&TypePasien{}).Where(&TypePasien{CodeName: codeName}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *TypePasienConn) GetTypePasienById(id int) (data *basemodel.TypePasienDataBaseModel, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&TypePasien{}).Where(&TypePasien{Id: id}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *TypePasienConn) TambahTypePasien(addData basemodel.TypePasienRequestInsertBaseModel) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Table("polis").
		Create(&addData)
	return result.Error
}

func (conn *TypePasienConn) EditTypePasien(editData basemodel.TypePasienRequestUpdateBaseModel, id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Model(&TypePasien{}).
		Where(&TypePasien{Id: id}).
		Updates(editData)
	return result.Error
}

func (conn *TypePasienConn) DeleteTypePasien(id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Delete(&TypePasien{Id: id})
	return result.Error
}
