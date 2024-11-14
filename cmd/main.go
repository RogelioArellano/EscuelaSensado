package main

import (
	"SensadoAlumnos/app"
	"log"
)

func main() {
	//Inicializar aplicación
	app, err := app.Initialize()
	if err != nil {
		log.Fatal("Error al inicializar microservicio:", err)
	}

	//Cerrar conexiones al finalizar la ejecución
	defer app.DB.Close()
	defer app.Service.KafkaProducer.Close()

	//Ejecutamos el proceso principal
	err = app.Service.ProcessStudents()
	if err != nil {
		log.Println("Error al procesar estudiantes")
	}

}
