package main

import (
	"SensadoAlumnos/app"
	"log"
)

// TODO mantener encendido el microservicio
func main() {
	//Inicializar aplicación
	app, err := app.Initialize()
	if err != nil {
		log.Fatal("Error al inicializar microservicio:", err)
	}

	//Cerrar conexiones al finalizar la ejecución
	defer app.DB.Close()
	defer app.Service.KafkaProducer.Close()

	log.Println("***********Iniciando procesamiento de estudiantes.***********")

	//Ejecutamos el proceso principal
	err = app.Service.ProcessStudents()
	if err != nil {
		log.Fatal("Error al procesar estudiantes")
	}

	log.Println("***********Procesamiento completado exitosamente.***********")

}
