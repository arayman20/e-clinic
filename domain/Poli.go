package domain

import (
	"fmt"
	"klinik/basemodel"
	"strings"
)

type PoliDomain struct {
	PoliRepo PoliInterface
}

func (dom *PoliDomain) DataPoli(pagination basemodel.PoliPaginationBaseModel) (code int, data []basemodel.PoliDataBaseModel, count int64, message string) {
	result, count, err := dom.PoliRepo.DataPoli(pagination)
	if result == nil && err == nil {
		return 404, nil, count, "Data not found!"
	}
	if err != nil {
		return 500, nil, count, "Internal Server Error"
	}

	for _, rowsData := range result {
		data = append(data, basemodel.PoliDataBaseModel{
			Id:          rowsData.Id,
			AntreanCode: rowsData.AntreanCode,
			CodeName:    rowsData.CodeName,
			Name:        rowsData.Name,
			Status:      rowsData.Status,
		})
	}
	return 200, data, count, "Data Poli"
}

func (dom *PoliDomain) TambahPoli(data basemodel.RequestInsertBaseModel) (code int, message string) {
	codeName := data.CodeName
	isOnlySpaces := codeName == ""
	if isOnlySpaces {
		return 400, "No Spece allowed"
	}

	isUppercase := codeName == strings.ToUpper(codeName)
	if !isUppercase {
		return 400, "Code Name need Uppercase letter"
	}
	result, err := dom.PoliRepo.GetPoliByCodeName(data.CodeName)
	if result != nil {
		return 400, "Code Name Already Exists"
	}
	if err != nil {
		return 500, err.Error()
	}
	insertErr := dom.PoliRepo.TambahPoli(data)
	if insertErr != nil {
		return 500, insertErr.Error()
	}
	return 201, fmt.Sprintf("Success add poli by code name %s", data.CodeName)
}

func (dom *PoliDomain) EditPoli(data basemodel.RequestUpdateBaseModel, id int) (code int, message string) {
	codeName := data.CodeName
	isOnlySpaces := codeName == ""
	if isOnlySpaces {
		return 400, "No Spece allowed"
	}

	isUppercase := codeName == strings.ToUpper(codeName)
	if !isUppercase {
		return 400, "Code Name need Uppercase letter"
	}
	result, err := dom.PoliRepo.GetPoliById(id)
	if result == nil && err == nil {
		return 400, "Failed Update Poli"
	}
	if err != nil {
		return 500, err.Error()
	}
	updateErr := dom.PoliRepo.EditPoli(data, id)
	if updateErr != nil {
		return 500, updateErr.Error()
	}
	return 201, fmt.Sprintf("Success edit poli by code name %s", result.CodeName)
}

func (dom *PoliDomain) DeletePoli(id int) (code int, message string) {
	result, err := dom.PoliRepo.GetPoliById(id)
	if result == nil && err == nil {
		return 400, "Failed Delete Poli"
	}
	deleteErr := dom.PoliRepo.DeletePoli(id)
	if deleteErr != nil {
		return 500, deleteErr.Error()
	}
	return 201, fmt.Sprintf("Success delete poli by code name %s", result.CodeName)
}
