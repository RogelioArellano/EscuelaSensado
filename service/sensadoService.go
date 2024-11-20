package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"

	"SensadoAlumnos/config"
	"SensadoAlumnos/models"
	"SensadoAlumnos/repository"
	"SensadoAlumnos/utils"
)

// Modelo del servicio de sensado
type SensadoService struct {
	EstudianteRepo  repository.EstudianteRepository
	InstruccionRepo repository.InstruccionEnvioRepository
	KafkaProducer   sarama.SyncProducer
}

// NewService crea una nueva instancia del servicio con sus dependencias
func NewService(estudianteRepo repository.EstudianteRepository, instruccionRepo repository.InstruccionEnvioRepository, kafkaProducer sarama.SyncProducer) *SensadoService {
	return &SensadoService{
		EstudianteRepo:  estudianteRepo,
		InstruccionRepo: instruccionRepo,
		KafkaProducer:   kafkaProducer,
	}
}

// ProcessStudents contiene la lógica principal del microservicio
func (s *SensadoService) ProcessStudents() error {
	//Leer estudiantes que estén dados de alta
	estudiantes, err := s.EstudianteRepo.LeerEstudiantesActivos()
	if err != nil {
		return nil
	}

	if len(estudiantes) == 0 {
		log.Println("No se encontraron estudiantes dados de alta")
		return nil
	}

	//Generamos la instruccion
	instruccion := models.InstruccionEnvio{
		FechaOperacion:     time.Now().Format("2006-01-02 15:04:05"),
		ClaveEmisor:        utils.ClaveEmisor,
		FolioConsecutivo:   s.obtenerFolioConsecutivo(),
		NumAltaEstudiantes: int32(len(estudiantes)),
		InfoAdicional:      estudiantes,
	}

	//Guardamos la instrucción de pago TODO agregar transaccionalidad
	err = s.InstruccionRepo.CreateInstruccion(instruccion)
	if err != nil {
		return err
	}

	//Convertimos a JSON
	instruccionJSON, err := json.Marshal(instruccion)
	if err != nil {
		return err
	}

	err = config.EnviarMensajeKafka(s.KafkaProducer, string(instruccionJSON), "instrucciónEnvio")
	if err != nil {
		return err
	}

	log.Println("Instrucción generada y enviada exitosamente")
	return nil
}

// obtenerFolioConsecutivo genera un folio consecutivo para la instrucción
func (s *SensadoService) obtenerFolioConsecutivo() int32 {
	lastFolio, err := s.InstruccionRepo.GetLastFolio()
	if err != nil {
		log.Println("Error al obtener el último folio:", err)
		return 1
	}
	return lastFolio + 1
}
