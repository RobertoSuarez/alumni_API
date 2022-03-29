package models

import "time"

type Empresa struct {
	ID               uint64 `json:"id" gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Nombre           string    `json:"nombre" gorm:"size:200"`
	Direccion        string    `json:"direccion" gorm:"size:500"`
	Correo           string    `json:"correo" gorm:"size:200"`
	SitioWeb         string    `json:"sitioweb" gorm:"size:200"`
	Ciudad           string    `json:"ciudad" gorm:"size:200"`
	Telefono1        string    `json:"telefono1" gorm:"size:200"`
	Telefono2        string    `json:"telefono2" gorm:"size:200"`
	Longitud         string    `json:"longitud" gorm:"size:200"`
	Latitud          string    `json:"latitud" gorm:"size:200"`
	UsuarioCreadorID uint64    `json:"usuarioCreadorID"`
	Usuarios         []Usuario `json:"usuarios" gorm:"many2many:usuario_empresas_asociadas;"`
}

func (Empresa) TableName() string {
	return "empresa"
}

func (e *Empresa) CrearEmpresa() error {
	tx := DB.Begin()

	err := tx.Create(&e).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// Listar todas las empresas registradas
func (Empresa) ObtenerEmpresas() (empresas []Empresa, err error) {
	empresas = []Empresa{}

	err = DB.Find(&empresas).Error
	if err != nil {
		return empresas, err
	}

	return empresas, nil
}