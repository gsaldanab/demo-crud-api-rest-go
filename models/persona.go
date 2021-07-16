package models

type Persona struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
	Telefono string `json:"telefono"`
}
