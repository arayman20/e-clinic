package domain

import (
	"fmt"
	"klinik/basemodel"
)

type PesertaBPJSDomain struct {
	PesertaBPJSRepo PesertaBPJSInterface
}

func (dom *PesertaBPJSDomain) DataPesertaBPJS(pagination basemodel.PesertaBPJSPaginationBaseModel) (code int, data []basemodel.PesertaBPJSWithMedicalRecordBaseModel, count int64, message string) {
	result, count, err := dom.PesertaBPJSRepo.DataPesertaBPJS(pagination)
	if result == nil && err == nil {
		return 404, nil, count, "Data not found!"
	}
	if err != nil {
		return 500, nil, count, "Internal Server Error"
	}

	for _, rowsData := range result {
		var dataMedical []basemodel.MedicalRecordBaseModel
		if rowsData.OneToManyMedicalRecord != nil {
			for _, dataMedicalRecord := range rowsData.OneToManyMedicalRecord {
				dataMedical = append(dataMedical, basemodel.MedicalRecordBaseModel{
					Id:           dataMedicalRecord.Id,
					DateRegister: dataMedicalRecord.DateRegister,
					TimeRegister: dataMedicalRecord.TimeRegister,
					PoliName:     dataMedicalRecord.OneToOnePoli.Name,
					Tension:      dataMedicalRecord.Tension,
					Weight:       dataMedicalRecord.Weight,
					Complain:     dataMedicalRecord.Complain,
					Handling:     dataMedicalRecord.Handling,
				})
			}

		}
		data = append(data, basemodel.PesertaBPJSWithMedicalRecordBaseModel{
			Id:                     rowsData.Id,
			NoBPJS:                 rowsData.NoBPJS,
			Name:                   rowsData.Name,
			BirthDate:              *rowsData.BirthDate,
			Address:                rowsData.Address,
			FaskesLevel:            rowsData.FaskesLevel,
			FaskesName:             rowsData.FaskesName,
			OneToManyMedicalRecord: dataMedical,
		})
	}
	return 200, data, count, "Data Peserta BPJS With Medical Record"
}

func (dom *PesertaBPJSDomain) DataPesertaBPJSMenyToMany() (code int, data []basemodel.PesertaBPJSWithMedicalRecordBaseModel, count int64, message string) {
	result, count, err := dom.PesertaBPJSRepo.DataPesertaBPJSMenyToMany()
	if result == nil && err == nil {
		return 404, nil, count, "Data not found!"
	}
	if err != nil {
		return 500, nil, count, "Internal Server Error"
	}

	for _, rowsData := range result {
		data = append(data, basemodel.PesertaBPJSWithMedicalRecordBaseModel{
			Id:          rowsData.Id,
			NoBPJS:      rowsData.NoBPJS,
			Name:        rowsData.Name,
			BirthDate:   *rowsData.BirthDate,
			Address:     rowsData.Address,
			FaskesLevel: rowsData.FaskesLevel,
			FaskesName:  rowsData.FaskesName,
		})
	}
	return 200, data, count, "Data Peserta BPJS With Medical Record"
}

func (dom *PesertaBPJSDomain) TambahPesertaBPJS(data basemodel.PesertaBPJSRequestInsertBaseModel) (code int, message string) {
	result, err := dom.PesertaBPJSRepo.GetPesertaBPJSByNoBPJS(data.NoBPJS)
	if result != nil {
		return 400, "No BPJS Already Exists"
	}
	if err != nil {
		return 500, err.Error()
	}
	insertErr := dom.PesertaBPJSRepo.TambahPesertaBPJS(data)
	if insertErr != nil {
		return 500, insertErr.Error()
	}
	return 201, fmt.Sprintf("Success add peserta by no bpjs %s", data.NoBPJS)
}

func (dom *PesertaBPJSDomain) EditPesertaBPJS(data basemodel.PesertaBPJSRequestUpdateBaseModel, id int) (code int, message string) {
	result, err := dom.PesertaBPJSRepo.GetPesertaBPJSById(id)
	if result == nil && err == nil {
		return 400, "Failed Update Peserta BPJS"
	}
	if err != nil {
		return 500, err.Error()
	}
	updateErr := dom.PesertaBPJSRepo.EditPesertaBPJS(data, id)
	if updateErr != nil {
		return 500, updateErr.Error()
	}
	return 201, fmt.Sprintf("Success edit peserta bpjs by no bpjs %s", result.NoBPJS)
}

func (dom *PesertaBPJSDomain) DeletePesertaBPJS(id int) (code int, message string) {
	result, err := dom.PesertaBPJSRepo.GetPesertaBPJSById(id)
	if result == nil && err == nil {
		return 400, "Failed Delete Peserta BPJS"
	}
	deleteErr := dom.PesertaBPJSRepo.DeletePesertaBPJS(id)
	if deleteErr != nil {
		return 500, deleteErr.Error()
	}
	return 201, fmt.Sprintf("Success delete peserta BPJS by no bpjs %s", result.NoBPJS)
}
