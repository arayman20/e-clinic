package domain

import (
	"fmt"
	"klinik/basemodel"
	"strings"
)

type TypePasienDomain struct {
	TypePasienRepo TypePasienInterface
}

func (dom *TypePasienDomain) DataTypePasien(pagination basemodel.TypePasienPaginationBaseModel) (code int, data []basemodel.TypePasienDataBaseModel, count int64, message string) {
	result, count, err := dom.TypePasienRepo.DataTypePasien(pagination)
	if result == nil && err == nil {
		return 404, nil, count, "Data not found!"
	}
	if err != nil {
		return 500, nil, count, "Internal Server Error"
	}

	for _, rowsData := range result {
		data = append(data, basemodel.TypePasienDataBaseModel{
			Id:       rowsData.Id,
			CodeName: rowsData.CodeName,
			Name:     rowsData.Name,
			Status:   rowsData.Status,
		})
	}
	return 200, data, count, "Data Tipe Pasien"
}

func (dom *TypePasienDomain) TambahTypePasien(data basemodel.TypePasienRequestInsertBaseModel) (code int, message string) {
	codeName := data.CodeName
	isOnlySpaces := codeName == ""
	if isOnlySpaces {
		return 400, "No Spece allowed"
	}

	isUppercase := codeName == strings.ToUpper(codeName)
	if !isUppercase {
		return 400, "Code Name need Uppercase letter"
	}
	result, err := dom.TypePasienRepo.GetTypePasienByCodeName(data.CodeName)
	if result != nil {
		return 400, "Code Name Already Exists"
	}
	if err != nil {
		return 500, err.Error()
	}
	insertErr := dom.TypePasienRepo.TambahTypePasien(data)
	if insertErr != nil {
		return 500, insertErr.Error()
	}
	return 201, fmt.Sprintf("Success add Tipe Pasien by code name %s", data.CodeName)
}

func (dom *TypePasienDomain) EditTypePasien(data basemodel.TypePasienRequestUpdateBaseModel, id int) (code int, message string) {
	codeName := data.CodeName
	isOnlySpaces := codeName == ""
	if isOnlySpaces {
		return 400, "No Spece allowed"
	}

	isUppercase := codeName == strings.ToUpper(codeName)
	if !isUppercase {
		return 400, "Code Name need Uppercase letter"
	}
	result, err := dom.TypePasienRepo.GetTypePasienById(id)
	if result == nil && err == nil {
		return 400, "Failed Update TypePasien"
	}
	if err != nil {
		return 500, err.Error()
	}
	updateErr := dom.TypePasienRepo.EditTypePasien(data, id)
	if updateErr != nil {
		return 500, updateErr.Error()
	}
	return 201, fmt.Sprintf("Success edit Tipe Pasien by code name %s", result.CodeName)
}

func (dom *TypePasienDomain) DeleteTypePasien(id int) (code int, message string) {
	result, err := dom.TypePasienRepo.GetTypePasienById(id)
	if result == nil && err == nil {
		return 400, "Failed Delete TypePasien"
	}
	deleteErr := dom.TypePasienRepo.DeleteTypePasien(id)
	if deleteErr != nil {
		return 500, deleteErr.Error()
	}
	return 201, fmt.Sprintf("Success delete Tipe Pasien by code name %s", result.CodeName)
}
