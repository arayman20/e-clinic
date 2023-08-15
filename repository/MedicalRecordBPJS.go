package repository

import (
	"errors"
	"klinik/basemodel"

	"gorm.io/gorm"
)

type MedicalRecordConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *MedicalRecordConn) DataMedicalRecord(pagination basemodel.MedicalRecordPaginationBaseModel) (data []basemodel.MedicalRecordBPJS, count int64, err error) {
	if conn.Err != nil {
		return nil, 0, conn.Err
	}

	db := conn.DB.Model(&basemodel.MedicalRecordBPJS{})
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	result := db.Debug().
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

func (conn *MedicalRecordConn) MedicalrecordGetPesertaBPJSByNoBPJS(noBpjs string) (data *basemodel.PesertaBPJS, err error) {
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

func (conn *MedicalRecordConn) MedicalrecordGetPesertaBPJSById(id int) (data *basemodel.PesertaBPJS, err error) {
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

func (conn *MedicalRecordConn) TambahMedicalRecord(addData basemodel.MedicalRecordRequestInsertBaseModel) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Table("medical_record_bpjs").
		Create(&addData)
	return result.Error
}

func (conn *MedicalRecordConn) EditMedicalRecord(editData basemodel.MedicalRecordRequestUpdateBaseModel, id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Model(&basemodel.MedicalRecordBPJS{}).
		Where(&basemodel.MedicalRecordBPJS{Id: id}).
		Updates(editData)
	return result.Error
}

func (conn *MedicalRecordConn) DeletePesertaBPJS(id int) (err error) {
	if conn.Err != nil {
		return conn.Err
	}
	result := conn.DB.Debug().Where(&basemodel.PesertaBPJS{Id: id}).Delete(&basemodel.PesertaBPJS{})
	return result.Error
}
