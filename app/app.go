package app

import (
	"SensadoAlumnos/config"
	"SensadoAlumnos/repository"
	"SensadoAlumnos/service"
	"database/sql"
)

type App struct {
	Service *service.SensadoService
	DB      *sql.DB
}

// Initialize configura las dependencias y así crear la instancia de Appp
func Initialize() (*App, error) {
	//Conectar a DB
	db, err := config.ConectDB()
	if err != nil {
		return nil, err
	}

	//Crear repos
	estudianteRepo := repository.NewEstudianteRepo(db)
	instruccionRepo := repository.NewInstruccionRepo(db)

	//Inicializar el productor de kafka
	kafkaProducer, err := config.InitKafkaProducer()
	if err != nil {
		return nil, err
	}

	svc := service.NewService(estudianteRepo, instruccionRepo, kafkaProducer)

	app := &App{
		Service: svc,
		DB:      db,
	}

	return app, nil
}
