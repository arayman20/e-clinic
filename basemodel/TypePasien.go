package basemodel

type TypePasienDataBaseModel struct {
	Id       int    `json:"id"`
	CodeName string `json:"codeName"`
	Name     string `json:"name"`
	Status   bool   `json:"status"`
}

type TypePasienRequestInsertBaseModel struct {
	CodeName string `json:"codeName"`
	Name     string `json:"name"`
	Status   bool   `json:"status"`
}

type TypePasienRequestUpdateBaseModel struct {
	CodeName string `json:"codeName"`
	Name     string `json:"name"`
	Status   bool   `json:"status"`
}

type TypePasienPaginationBaseModel struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
