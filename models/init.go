package models

import (
	"github.com/RobertoSuarez/apialumni/database"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBaseTable(config *viper.Viper) {

	DB = database.GetInstancia(config)

	// Migraciones
	DB.AutoMigrate(&Usuario{})
	DB.AutoMigrate(&Empleo{})
	DB.AutoMigrate(&Educacion{})
	DB.AutoMigrate(&RoleUsuario{})
	DB.AutoMigrate(&Grupo{})
	DB.AutoMigrate(&Trabajo{})
	DB.AutoMigrate(&Empresa{})
	DB.AutoMigrate(&Area{})
	DB.AutoMigrate(&Subarea{})
	DB.AutoMigrate(&Provincia{})
	DB.AutoMigrate(&Ciudad{})
	DB.AutoMigrate(&LogBusquedas{})
}
