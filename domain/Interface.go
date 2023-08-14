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
	GetPesertaByNoBPJS(noBpjs string) (data *PesertaBPJS, err error)
	RegisterPesertaBPJS(dataRegister PesertaBPJS) (err error)
}

type AntreanBPJSInterface interface {
	AntreanGetPesertaByNoBPJS(noBpjs string) (data PesertaBPJS, err error)
	InsertPendaftaran(dataRegister AntreanBPJS) (err error)
	CallPendaftaran() (noAntrean string, err error)
	DeactiveStatusAntrean(antreanId int) (err error)
	GetLastNumberAntrean(poliId int) (antrean int, err error)
}

type PoliInterface interface {
	DataPoli(pagination basemodel.PoliPaginationBaseModel) (data []basemodel.PoliDataBaseModel, count int64, err error)
	GetPoliByCodeName(codeName string) (data *basemodel.PoliDataBaseModel, err error)
	GetPoliById(id int) (data *basemodel.PoliDataBaseModel, err error)
	TambahPoli(data basemodel.RequestInsertBaseModel) (err error)
	EditPoli(editData basemodel.RequestUpdateBaseModel, id int) (err error)
	DeletePoli(id int) (err error)
}

type TypePasienInterface interface {
	DataTypePasien(pagination basemodel.TypePasienPaginationBaseModel) (data []basemodel.TypePasienDataBaseModel, count int64, err error)
	GetTypePasienByCodeName(codeName string) (data *basemodel.TypePasienDataBaseModel, err error)
	GetTypePasienById(id int) (data *basemodel.TypePasienDataBaseModel, err error)
	TambahTypePasien(data basemodel.TypePasienRequestInsertBaseModel) (err error)
	EditTypePasien(editData basemodel.TypePasienRequestUpdateBaseModel, id int) (err error)
	DeleteTypePasien(id int) (err error)
}
