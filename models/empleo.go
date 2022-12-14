package models

import (
	"errors"
	"time"
)

// Esto es el modelo de empleo, el cual son las ofertas de empleos
// que las empresas publican

type Empleo struct {
	ID                       uint64 `json:"id" gorm:"primary_key"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
	Titulo                   string    `json:"titulo" gorm:"size:200"`
	Descripcion              string    `json:"descripcion"`
	Profesion                string    `json:"profesion" gorm:"size:200"`
	Puesto                   string    `json:"puesto" gorm:"size:200"`
	TipoEmplo                string    `json:"tipoEmpleo" gorm:"size:200"` //Modalidad de trabajo
	AreaID                   uint64    `json:"areaid,omitempty"`           // se realaciona con la tabla area
	Area                     Area      `json:"area" gorm:"foreignKey:AreaID"`
	SubareaID                uint64    `json:"subareaid,omitempty"` // se relaciona con la tabla subarea
	Subarea                  Subarea   `json:"subarea" gorm:"foreignKey:SubareaID"`
	Sueldo                   string    `json:"sueldo" gorm:"size:200"`
	TiempoExperiencia        string    `json:"tiempoExperiencia" gorm:"size:200"` // Los años de experiencia
	Jornada                  string    `json:"jornada" gorm:"size:200"`
	TipoContrato             string    `json:"tipoContrato" gorm:"size:200"`
	ConocimientosAdicionales string    `json:"conocimientosAdicionales" gorm:"size:200"`
	ProvinciaID              uint64    `json:"provinciaid"`
	Provincia                Provincia `json:"provincia" gorm:"foreignKey:ProvinciaID"`
	CiudadID                 uint64    `json:"ciudadid"`
	Ciudad                   Ciudad    `json:"ciudad" gorm:"foreignKey:CiudadID"`
	PostulanteDiscapacidad   *bool     `json:"postulanteDiscapacidad" gorm:"default:false"` // si el trabajo es para personas con capacidades limitadas.
	Publicado                time.Time `json:"publicado"`
	Borrador                 *bool     `json:"borrador" gorm:"default:false"`
	EmpresaID                uint64    `json:"empresaid"`
	UsuarioID                uint64    `json:"usuario_id"`
	Usuario                  Usuario   `json:"usuario" gorm:"foreignKey:UsuarioID"`
	Activo                   *bool     `json:"activo"`                                                // el estado activo representa si, se aceptan aplicaciones al trabajo, o no
	UsuariosAplicados        []Usuario `json:"usuariosAplicados" gorm:"many2many:empleos_aplicados;"` // Usuarios que han aplicado a este trabajo.
	//Area                     string    `json:"area" gorm:"size:200"`       // Categoria
}

func (Empleo) TableName() string {
	return "empleo"
}

// Hacer: Publicar empleo
func (e *Empleo) Crear() error {
	if e.EmpresaID < 1 {
		return errors.New("fatal no existe el id de la empresa")
	}

	e.Publicado = time.Now()

	activo := true
	e.Activo = &activo

	result := DB.Create(&e)
	if result.Error != nil {
		return result.Error
	}

	DB.Model(&e).Preload("Area").Preload("Subarea").First(&e)

	return nil
}

// Actulizar empleo
func (e *Empleo) Actualizar() error {
	tx := DB.Begin()

	result := tx.Model(&e).Updates(Empleo{
		Titulo:                   e.Titulo,
		Descripcion:              e.Descripcion,
		Profesion:                e.Profesion,
		Puesto:                   e.Puesto,
		TipoEmplo:                e.TipoEmplo,
		AreaID:                   e.AreaID,
		SubareaID:                e.SubareaID,
		Sueldo:                   e.Sueldo,
		TiempoExperiencia:        e.TiempoExperiencia,
		Jornada:                  e.Jornada,
		TipoContrato:             e.TipoContrato,
		ConocimientosAdicionales: e.ConocimientosAdicionales,
		Ciudad:                   e.Ciudad,
		PostulanteDiscapacidad:   e.PostulanteDiscapacidad,
		Borrador:                 e.Borrador,
		//Area:                     e.Area,
	})

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Preload("Area").Preload("Subarea").First(&e, e.ID)
	e.AreaID = 0
	e.SubareaID = 0

	tx.Commit()
	return nil
}

// Cambiar estado del empleo

// Listar los empleos
// Fala buscar por palabras claves y ciudad
func (Empleo) ObtenerTodos(offset int, pageSize int, maps interface{}, title string) (empleos []Empleo, err error) {

	result := DB.
		Where("titulo LIKE ?", "%"+title+"%").
		Where(maps).
		Preload("Area").
		Preload("Subarea").
		Preload("Provincia").
		Preload("Ciudad").
		Order("created_at desc").
		Offset(offset).Limit(pageSize).Find(&empleos)
	if result.Error != nil {
		return empleos, result.Error
	}

	return empleos, nil
}

// Obtener empleo por el id
func (e *Empleo) ObtenerEmpleoByID() error {

	result := DB.Model(&e).Preload("Area").Preload("Subarea").First(&e)
	if result.Error != nil {
		return result.Error
	}
	e.AreaID = 0
	e.SubareaID = 0

	return nil
}

// cambiar el estado de un empleo, que si esta activo o inactivo
func (e *Empleo) CambiarEstado() (Empleo, error) {

	tx := DB.Begin()

	result := tx.Model(&e).Updates(Empleo{
		Activo: e.Activo,
	})

	if result.Error != nil {
		tx.Rollback()
		return *e, result.Error
	}

	tx.First(&e)

	tx.Commit()
	return *e, nil
}
