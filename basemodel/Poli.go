package basemodel

type Poli struct {
	Id          int    `gorm:"type:bigint;primaryKey;autoIncrement;index"`
	AntreanCode string `gorm:"type:varchar"`
	CodeName    string `gorm:"varchar"`
	Name        string `gorm:"varchar"`
	Status      bool   `gorm:"type:boolean;default:True"`
}

type PoliDataBaseModel struct {
	Id          int    `json:"id"`
	AntreanCode string `json:"antreanCode"`
	CodeName    string `json:"codeName"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
}

type RequestInsertBaseModel struct {
	AntreanCode string `json:"antreanCode"`
	CodeName    string `json:"codeName"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
}

type RequestUpdateBaseModel struct {
	AntreanCode string `json:"antreanCode"`
	CodeName    string `json:"codeName"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
}

type PoliPaginationBaseModel struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
