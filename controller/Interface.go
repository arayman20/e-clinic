package controller

import "klinik/basemodel"

type PoliControllerInterface interface {
	DataPoli(pagination basemodel.PoliPaginationBaseModel) (code int, data []basemodel.PoliDataBaseModel, count int64, message string)
	TambahPoli(data basemodel.RequestInsertBaseModel) (code int, message string)
	EditPoli(data basemodel.RequestUpdateBaseModel, id int) (code int, message string)
	DeletePoli(id int) (code int, message string)
}
