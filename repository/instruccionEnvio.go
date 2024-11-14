package repository

import (
	"SensadoAlumnos/models"
	"database/sql"
)

// Interfaz para el repo de Instrucciones de envio
type InstruccionEnvioRepository interface {
	CreateInstruccion(instruccion models.InstruccionEnvio) error
}

// Estructura que implementa la interfaz de Instrucciones de envio
type instruccionRepo struct {
	db *sql.DB
}

// Constructor para el repo de Instrucciones de envio
func NewInstruccionRepo(db *sql.DB) InstruccionEnvioRepository {
	return &instruccionRepo{db: db}
}

func (i *instruccionRepo) CreateInstruccion(instruccion models.InstruccionEnvio) error {
	_, err := i.db.Exec("INSERT INTO instrucciones (fechaOperacion, claveEmisor, folioConsecutivo, numAltaEstudiantes) VALUES ($1, $2, $3, $4)",
		instruccion.FechaOperacion, instruccion.ClaveEmisor, instruccion.FolioConsecutivo, instruccion.NumAltaEstudiantes)
	return err
}
