package repository

import (
	"errors"
	"fmt"
	"klinik/basemodel"

	"gorm.io/gorm"
)

type PesertaBPJSConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *PesertaBPJSConn) DataPesertaBPJS(pagination basemodel.PesertaBPJSPaginationBaseModel) (data []basemodel.PesertaBPJS, count int64, err error) {
	// one to many query dan one to one
	if conn.Err != nil {
		return nil, 0, conn.Err
	}

	db := conn.DB.Preload("OneToManyMedicalRecord").Model(&basemodel.PesertaBPJS{})
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
	if len(data) == 0 {
		return nil, 0, nil
	}
	return data, count, nil

}

func (conn *PesertaBPJSConn) DataPesertaBPJSMenyToMany() (data []basemodel.PesertaBPJS, count int64, err error) {
	if conn.Err != nil {
		return nil, 0, conn.Err
	}

	db := conn.DB.Preload("ManyToManyRecord").Model(&basemodel.PesertaBPJS{})
	result := db.
		Find(&data)
	fmt.Println(data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, result.Error
	}
	if len(data) == 0 {
		return nil, 0, nil
	}
	return data, count, nil

}

func (conn *PesertaBPJSConn) GetPesertaBPJSByNoBPJS(noBpjs string) (data *basemodel.PesertaBPJS, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&basemodel.PesertaBPJS{}).Where(&basemodel.PesertaBPJS{NoBPJS: noBpjs}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *PesertaBPJSConn) GetPesertaBPJSById(id int) (data *basemodel.PesertaBPJS, err error) {
	if conn.Err != nil {
		return nil, conn.Err
	}

	result := conn.DB.Model(&basemodel.PesertaBPJS{}).Where(&basemodel.PesertaBPJS{Id: id}).First(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return data, nil
}

func (conn *PesertaBPJSConn) TambahPesertaBPJS(addData basemodel.PesertaBPJSRequestInsertBaseModel) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Table("peserta_bpjs").
		Create(&addData)
	return result.Error
}

func (conn *PesertaBPJSConn) EditPesertaBPJS(editData basemodel.PesertaBPJSRequestUpdateBaseModel, id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Model(&basemodel.PesertaBPJS{}).
		Where(&basemodel.PesertaBPJS{Id: id}).
		Updates(editData)
	return result.Error
}

func (conn *PesertaBPJSConn) DeletePesertaBPJS(id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Where(&basemodel.PesertaBPJS{Id: id}).Delete(&basemodel.PesertaBPJS{})
	return result.Error
}
