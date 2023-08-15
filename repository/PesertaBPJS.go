package repository

import (
	"errors"
	"klinik/basemodel"

	"gorm.io/gorm"
)

type PesertaBPJSConn struct {
	DB  *gorm.DB
	Err error
}

func (conn *PesertaBPJSConn) DataPesertaBPJS(pagination basemodel.PesertaBPJSPaginationBaseModel) (data []basemodel.PesertaBPJS, count int64, err error) {
	if conn.Err != nil {
		return nil, 0, conn.Err
	}

	db := conn.DB.Preload("OneToManyMedicalRecord").Model(&basemodel.PesertaBPJS{})
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

	// for _, dataPeserta := range dataRows {
	// 	var medicalRecord []basemodel.MedicalRecordBaseModel
	// 	if dataPeserta.OneToManyMedicalRecord != nil {
	// 		for _, dataMedical := range dataPeserta.OneToManyMedicalRecord {
	// 			medicalRecord = append(medicalRecord, basemodel.MedicalRecordBaseModel{
	// 				Id:           dataMedical.Id,
	// 				NoBPJS:       dataMedical.NoBPJS,
	// 				DateRegister: dataMedical.DateRegister,
	// 				TimeRegister: dataMedical.TimeRegister,
	// 				PoliId:       dataMedical.PoliId,
	// 				PoliName:     dataMedical.OneToOnePoli.Name,
	// 				Tension:      dataMedical.Tension,
	// 				Weight:       dataMedical.Weight,
	// 				Complain:     dataMedical.Complain,
	// 				Handling:     dataMedical.Handling,
	// 			})
	// 		}
	// 		data = append(data, basemodel.PesertaBPJSWithMedicalRecordBaseModel{
	// 			Id:                     dataPeserta.Id,
	// 			NoBPJS:                 dataPeserta.NoBPJS,
	// 			Name:                   dataPeserta.Name,
	// 			BirthDate:              *dataPeserta.BirthDate,
	// 			Address:                dataPeserta.Address,
	// 			FaskesLevel:            dataPeserta.FaskesLevel,
	// 			FaskesName:             dataPeserta.FaskesName,
	// 			OneToManyMedicalRecord: medicalRecord,
	// 		})

	// 	}
	// }

	return data, count, nil

}

func (conn *PesertaBPJSConn) GetPesertaBPJSByNoBPJS(noBpjs int) (data *basemodel.PesertaBPJS, err error) {
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
	result := conn.DB.Debug().Where(&basemodel.PesertaBPJS{Id: id}).Delete(&basemodel.PesertaBPJS{})
	return result.Error
}
