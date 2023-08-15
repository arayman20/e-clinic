package repository

import (
	"errors"
	"klinik/basemodel"

	"gorm.io/gorm"
)

type PoliConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *PoliConn) DataPoli(pagination basemodel.PoliPaginationBaseModel) (data []basemodel.Poli, count int64, err error) {
	if conn.Err != nil {
		return nil, 0, conn.Err
	}

	db := conn.DB.Model(&basemodel.Poli{})
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

func (conn *PoliConn) GetPoliByCodeName(codeName string) (data *basemodel.Poli, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&basemodel.Poli{}).Where(&basemodel.Poli{CodeName: codeName}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *PoliConn) GetPoliById(id int) (data *basemodel.Poli, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&basemodel.Poli{}).Where(&basemodel.Poli{Id: id}).First(&data)
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
	result := conn.DB.Model(&basemodel.Poli{}).
		Where(&basemodel.Poli{Id: id}).
		Updates(editData)
	return result.Error
}

func (conn *PoliConn) DeletePoli(id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Delete(&basemodel.Poli{Id: id})
	return result.Error
}
