package repository

import (
	"SensadoAlumnos/models"
	"database/sql"
)

type EstudianteRepository interface {
	GetEstudiantesActivos() ([]models.Estudiante, error)
}

//Estructura para implementar
type EstudianteRepo struct {
	db *sql.DB
}

//Constructor para el repositorio de estudiantes
func NewEstudianteRepo(db *sql.DB) EstudianteRepository {
	return &EstudianteRepo{db: db}
}

func (r *EstudianteRepo) GetEstudiantesActivos() ([]models.Estudiante, error) {
	rows, err := r.db.Query("SELECT idEstudiante, nombre, direccion, email, telefono, altaLocal, altaSep FROM estudiantes WHERE altaLocal = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var estudiantes []models.Estudiante
	for rows.Next() {
		var estudiante models.Estudiante
		if err := rows.Scan(&estudiante.IDEstudiante, &estudiante.Nombre, &estudiante.Direccion, &estudiante.Email, &estudiante.Telefono, &estudiante.AltaLocal, &estudiante.AltaSep); err != nil {
			return nil, err
		}
		estudiantes = append(estudiantes, estudiante)
	}
	return estudiantes, nil
}
