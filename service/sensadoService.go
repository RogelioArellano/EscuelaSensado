package service

import (
	"SensadoAlumnos/models"
	"SensadoAlumnos/repository"
	"SensadoAlumnos/utils"
	"time"
)

func GenerarInstruccionEnvio(repo repository.EstudianteRepository) (models.InstruccionEnvio, error) {
    estudiantesActivos, err := repo.GetEstudiantesActivos()
    if err != nil {
        return models.InstruccionEnvio{}, err
    }

    instruccion := models.InstruccionEnvio{
        InstruccionID:      1,  // Esto podría ser un ID autogenerado o secuencial
        FechaOperacion:     time.Now().Format("2006-01-02"),
        ClaveEmisor:        utils.ClaveEmisor, // Ajusta el valor según sea necesario
        FolioConsecutivo:   1,    // Puede ser un contador o valor secuencial
        NumAltaEstudiantes: int32(len(estudiantesActivos)),
        InfoAdicional:      estudiantesActivos,
    }

    return instruccion, nil
}