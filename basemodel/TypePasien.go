package basemodel

type TypePasien struct {
	Id       int    `gorm:"type:bigint;primaryKey;autoIncrement;index"`
	CodeName string `gorm:"varchar"`
	Name     string `gorm:"varchar"`
	Status   bool   `gorm:"type:boolean;default:True"`
}

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
