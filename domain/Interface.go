package domain

import (
	"klinik/basemodel"
	"time"
)

type PesertaBPJS struct {
	Id          int
	NoBPJS      string
	Name        string
	BirthDate   time.Time
	Address     string
	FaskesLevel string
	FaskesName  string
}

type AntreanBPJS struct {
	Id           int
	NoBPJS       string
	DateRegister time.Time
	TimeRegister time.Time
	PoliId       int
	NoAntrean    int
	TypePasien   int
	Status       bool
}

type PaginationPoli struct {
	Page int
	Size int
}

type PesertaBPJSInterface interface {
	DataPesertaBPJS(pagination basemodel.PesertaBPJSPaginationBaseModel) (data []basemodel.PesertaBPJS, count int64, err error)
	GetPesertaBPJSByNoBPJS(noBpjs string) (data *basemodel.PesertaBPJS, err error)
	GetPesertaBPJSById(id int) (data *basemodel.PesertaBPJS, err error)
	TambahPesertaBPJS(addData basemodel.PesertaBPJSRequestInsertBaseModel) (err error)
	EditPesertaBPJS(editData basemodel.PesertaBPJSRequestUpdateBaseModel, id int) (err error)
	DeletePesertaBPJS(id int) (err error)
}

type PoliInterface interface {
	DataPoli(pagination basemodel.PoliPaginationBaseModel) (data []basemodel.Poli, count int64, err error)
	GetPoliByCodeName(codeName string) (data *basemodel.Poli, err error)
	GetPoliById(id int) (data *basemodel.Poli, err error)
	TambahPoli(data basemodel.RequestInsertBaseModel) (err error)
	EditPoli(editData basemodel.RequestUpdateBaseModel, id int) (err error)
	DeletePoli(id int) (err error)
}

type TypePasienInterface interface {
	DataTypePasien(pagination basemodel.TypePasienPaginationBaseModel) (data []basemodel.TypePasien, count int64, err error)
	GetTypePasienByCodeName(codeName string) (data *basemodel.TypePasien, err error)
	GetTypePasienById(id int) (data *basemodel.TypePasien, err error)
	TambahTypePasien(data basemodel.TypePasienRequestInsertBaseModel) (err error)
	EditTypePasien(editData basemodel.TypePasienRequestUpdateBaseModel, id int) (err error)
	DeleteTypePasien(id int) (err error)
}

type MedicalRecordInterface interface {
	DataMedicalRecord(pagination basemodel.MedicalRecordPaginationBaseModel) (data []basemodel.MedicalRecordBPJS, count int64, err error)
	MedicalrecordGetPesertaBPJSByNoBPJS(noBpjs string) (data *basemodel.PesertaBPJS, err error)
	MedicalrecordGetPesertaBPJSById(id int) (data *basemodel.PesertaBPJS, err error)
	TambahMedicalRecord(addData basemodel.MedicalRecordRequestInsertBaseModel) (err error)
	EditMedicalRecord(editData basemodel.MedicalRecordRequestUpdateBaseModel, id int) (err error)
	DeletePesertaBPJS(id int) (err error)
}
