package basemodel

import "time"

type MedicalRecordBPJS struct {
	Id           int       `gorm:"type:bigint;autoIncrement;index"`
	NoBPJS       int       `gorm:"type:bigint"`
	DateRegister time.Time `gorm:"type:date;default:null;default:CURRENT_DATE;index"`
	TimeRegister time.Time `gorm:"type:timestamp(0);default:CURRENT_TIMESTAMP"`
	PoliId       int       `gorm:"type:integer;index"`
	OneToOnePoli Poli      `gorm:"foreignKey:Id"`
	Tension      string    `gorm:"type:varchar"`
	Weight       int       `gorm:"type:integer"`
	Complain     string    `gorm:"type:text"`
	Handling     string    `gorm:"type:text"`
}

type MedicalRecordBaseModel struct {
	Id           int       `json:"id"`
	NoBPJS       int       `json:"noBPJS"`
	DateRegister time.Time `json:"dateRegister"`
	TimeRegister time.Time `json:"timeRegister"`
	PoliId       int       `json:"poliId"`
	PoliName     string    `json:"poliName"`
	Tension      string    `json:"tension"`
	Weight       int       `json:"weight"`
	Complain     string    `json:"complain"`
	Handling     string    `json:"handling"`
}
