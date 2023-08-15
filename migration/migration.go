package migration

import (
	"klinik/basemodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Migrator().DropTable(&basemodel.TypePasien{},
		&basemodel.Poli{},
		&basemodel.PesertaBPJS{},
		&basemodel.MedicalRecordBPJS{})

	db.AutoMigrate(&basemodel.TypePasien{})
	db.AutoMigrate(&basemodel.Poli{})
	db.AutoMigrate(&basemodel.PesertaBPJS{})
	db.AutoMigrate(&basemodel.MedicalRecordBPJS{})

}
