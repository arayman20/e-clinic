package basemodel

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
