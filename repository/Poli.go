package repository

import (
	"errors"
	"klinik/basemodel"

	"gorm.io/gorm"
)

type Poli struct {
	Id          int    `gorm:"type:bigint;primaryKey;autoIncrement;index"`
	AntreanCode string `gorm:"type:varchar"`
	CodeName    string `gorm:"varchar"`
	Name        string `gorm:"varchar"`
	Status      bool   `gorm:"type:boolean;default:True"`
}

type PoliConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *PoliConn) DataPoli(pagination basemodel.PoliPaginationBaseModel) (data []basemodel.PoliDataBaseModel, count int64, err error) {
	if conn.Err != nil {
		return nil, 0, conn.Err
	}

	db := conn.DB.Model(&Poli{})
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

func (conn *PoliConn) GetPoliByCodeName(codeName string) (data *basemodel.PoliDataBaseModel, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&Poli{}).Where(&Poli{CodeName: codeName}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *PoliConn) GetPoliById(id int) (data *basemodel.PoliDataBaseModel, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&Poli{}).Where(&Poli{Id: id}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *PoliConn) TambahPoli(addData basemodel.RequestInsertBaseModel) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Table("polis").
		Create(&addData)
	return result.Error
}

func (conn *PoliConn) EditPoli(editData basemodel.RequestUpdateBaseModel, id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Model(&Poli{}).
		Where(&Poli{Id: id}).
		Updates(editData)
	return result.Error
}

func (conn *PoliConn) DeletePoli(id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Delete(&Poli{Id: id})
	return result.Error
}
