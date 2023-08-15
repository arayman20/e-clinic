package controller

import "klinik/basemodel"

type PoliControllerInterface interface {
	DataPoli(pagination basemodel.PoliPaginationBaseModel) (code int, data []basemodel.PoliDataBaseModel, count int64, message string)
	TambahPoli(data basemodel.RequestInsertBaseModel) (code int, message string)
	EditPoli(data basemodel.RequestUpdateBaseModel, id int) (code int, message string)
	DeletePoli(id int) (code int, message string)
}

type TypePasienControllerInterface interface {
	DataTypePasien(pagination basemodel.TypePasienPaginationBaseModel) (code int, data []basemodel.TypePasienDataBaseModel, count int64, message string)
	TambahTypePasien(data basemodel.TypePasienRequestInsertBaseModel) (code int, message string)
	EditTypePasien(data basemodel.TypePasienRequestUpdateBaseModel, id int) (code int, message string)
	DeleteTypePasien(id int) (code int, message string)
}

type PesertaBPJSControllerInterface interface {
	DataPesertaBPJS(pagination basemodel.PesertaBPJSPaginationBaseModel) (code int, data []basemodel.PesertaBPJSWithMedicalRecordBaseModel, count int64, message string)
	TambahPesertaBPJS(data basemodel.PesertaBPJSRequestInsertBaseModel) (code int, message string)
	EditPesertaBPJS(data basemodel.PesertaBPJSRequestUpdateBaseModel, id int) (code int, message string)
	DeletePesertaBPJS(id int) (code int, message string)
}
