package basemodel

import "time"

type MedicalRecordBPJS struct {
	Id                int            `gorm:"type:bigint;autoIncrement;index"`
	PesertaId         int            `gorm:"type:integer"`
	DateRegister      time.Time      `gorm:"type:date;default:null;default:CURRENT_DATE;index"`
	TimeRegister      time.Time      `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
	PoliId            int            `gorm:"type:integer;index"`
	OneToOnePoli      Poli           `gorm:"foreignKey:Id"`
	Tension           string         `gorm:"type:varchar"`
	Weight            int            `gorm:"type:integer"`
	Complain          string         `gorm:"type:text"`
	Handling          string         `gorm:"type:text"`
	ManyToManyPeserta []*PesertaBPJS `gorm:"many2many:peserta"`
}

type MedicalRecordBaseModel struct {
	Id           int       `json:"id"`
	PesertaId    int       `json:"pesertaId"`
	DateRegister time.Time `json:"dateRegister"`
	TimeRegister time.Time `json:"timeRegister"`
	PoliName     string    `json:"poliName"`
	Tension      string    `json:"tension"`
	Weight       int       `json:"weight"`
	Complain     string    `json:"complain"`
	Handling     string    `json:"handling"`
}

type MedicalRecordPaginationBaseModel struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type MedicalRecordRequestInsertBaseModel struct {
	PesertaId int    `json:"pesertaId"`
	PoliId    int    `json:"poliId"`
	Tension   string `json:"tension"`
	Weight    int    `json:"weight"`
	Complain  string `json:"complain"`
	Handling  string `json:"handling"`
}

type MedicalRecordRequestUpdateBaseModel struct {
	PesertaId int    `json:"pesertaId"`
	PoliId    int    `json:"poliId"`
	Tension   string `json:"tension"`
	Weight    int    `json:"weight"`
	Complain  string `json:"complain"`
	Handling  string `json:"handling"`
}
