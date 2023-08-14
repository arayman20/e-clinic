package migration

import (
	"klinik/repository"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&repository.Poli{})
	db.AutoMigrate(&repository.TypePasien{})
	db.AutoMigrate(&repository.PesertaBPJS{})
	db.AutoMigrate(&repository.AntreanBPJS{})

}
