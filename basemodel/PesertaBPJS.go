package basemodel

import "time"

type PesertaBPJS struct {
	Id                     int                  `gorm:"type:bigint;primaryKey;autoIncrement;index"`
	OneToManyMedicalRecord []MedicalRecordBPJS  `gorm:"foreignKey:PesertaId"`
	NoBPJS                 string               `gorm:"type:varchar;index"`
	Name                   string               `gorm:"type:varchar;index"`
	BirthDate              *time.Time           `gorm:"type:date"`
	Address                string               `gorm:"type:varchar"`
	FaskesLevel            string               `gorm:"type:varchar"`
	FaskesName             string               `gorm:"type:varchar"`
	ManyToManyRecord       []*MedicalRecordBPJS `gorm:"many2many:peserta"`
}

type PesertaBPJSDataBaseModel struct {
	Id          int       `json:"id"`
	NoBPJS      string    `json:"noBPJS"`
	Name        string    `json:"name"`
	BirthDate   time.Time `json:"birthDate"`
	Address     string    `json:"address"`
	FaskesLevel string    `json:"faskesLevel"`
	FaskesName  string    `json:"faskesName"`
}

type PesertaBPJSWithMedicalRecordBaseModel struct {
	Id                     int         `json:"id"`
	NoBPJS                 string      `json:"noBPJS"`
	Name                   string      `json:"name"`
	BirthDate              time.Time   `json:"birthDate"`
	Address                string      `json:"address"`
	FaskesLevel            string      `json:"faskesLevel"`
	FaskesName             string      `json:"faskesName"`
	OneToManyMedicalRecord interface{} `json:"medicalRecord"`
}

type PesertaBPJSPaginationBaseModel struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type PesertaBPJSRequestInsertBaseModel struct {
	NoBPJS      string `json:"noBPJS"`
	Name        string `json:"name"`
	BirthDate   string `json:"birthDate"`
	Address     string `json:"address"`
	FaskesLevel string `json:"faskesLevel"`
	FaskesName  string `json:"faskesName"`
}

type PesertaBPJSRequestUpdateBaseModel struct {
	NoBPJS      string `json:"noBPJS"`
	Name        string `json:"name"`
	BirthDate   string `json:"birthDate"`
	Address     string `json:"address"`
	FaskesLevel string `json:"faskesLevel"`
	FaskesName  string `json:"faskesName"`
}
