package domain

import (
	"fmt"
	"klinik/basemodel"
)

type MedicalRecordDomain struct {
	MedicalRecordRepo MedicalRecordInterface
}

func (dom *MedicalRecordDomain) DataPesertaBPJS(pagination basemodel.MedicalRecordPaginationBaseModel) (code int, data []basemodel.MedicalRecordBaseModel, count int64, message string) {
	result, count, err := dom.MedicalRecordRepo.DataMedicalRecord(pagination)
	if result == nil && err == nil {
		return 404, nil, count, "Data not found!"
	}
	if err != nil {
		return 500, nil, count, "Internal Server Error"
	}

	for _, rowsData := range result {

		data = append(data, basemodel.MedicalRecordBaseModel{
			Id:           rowsData.Id,
			PesertaId:    rowsData.PesertaId,
			DateRegister: rowsData.DateRegister,
			TimeRegister: rowsData.TimeRegister,
			PoliName:     rowsData.OneToOnePoli.Name,
			Tension:      rowsData.Tension,
			Weight:       rowsData.Weight,
			Complain:     rowsData.Complain,
			Handling:     rowsData.Handling,
		})
	}
	return 200, data, count, "Data Rekam medis"
}

func (dom *MedicalRecordDomain) TambahMedicalRecord(data basemodel.MedicalRecordRequestInsertBaseModel) (code int, message string) {
	insertErr := dom.MedicalRecordRepo.TambahMedicalRecord(data)
	if insertErr != nil {
		return 500, insertErr.Error()
	}
	return 201, fmt.Sprintf("Success add medical record by Id %s", data.PesertaId)
}

func (dom *MedicalRecordDomain) EditMedicalRecord(data basemodel.MedicalRecordRequestUpdateBaseModel, id int) (code int, message string) {
	updateErr := dom.MedicalRecordRepo.EditMedicalRecord(data, id)
	if updateErr != nil {
		return 500, updateErr.Error()
	}
	return 201, fmt.Sprintf("Success edit peserta bpjs by id %s", id)
}

func (dom *MedicalRecordDomain) DeleteMedicalRecord(id int) (code int, message string) {
	deleteErr := dom.MedicalRecordRepo.DeletePesertaBPJS(id)
	if deleteErr != nil {
		return 500, deleteErr.Error()
	}
	return 201, fmt.Sprintf("Success delete peserta BPJS by id %s", id)
}
