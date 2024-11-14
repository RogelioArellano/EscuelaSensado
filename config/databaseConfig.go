package config

import (
	"SensadoAlumnos/env"
	"database/sql"
	"fmt"
	"log"
)

func ConectDB() (*sql.DB, error) {
	host := env.GetEnv("DBSERVER", "localhost")
	port := env.GetEnv("DBPORT", "5432")
	user := env.GetEnv("DBUSER", "postgres")
	pass := env.GetEnv("DBPASSWORD", "password")
	dbname := env.GetEnv("DBNAME", "escuela")

	//Conectar a la base de datos (En este punto ya debería estar creada "escuela")
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL", err)
	}

	defer db.Close()

	//Probar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("Error conectando a la base de datos")
	}

	//Crear la tabla instrucción Envío
	crearInstruccionEnvio := `
	CREATE TABLE IF NOT EXISTS instrucciones(
		idInstruccion SERIAL PRIMARY KEY,
		fechaOperacion VARCHAR(20),
		claveEmisor INTEGER,
		folioConsecutivo INTEGER
	)`
	_, err = db.Exec(crearInstruccionEnvio)
	if err != nil {
		log.Fatal("Error al crear la tabla estudiantes", err)
	}

	return db, nil
}
