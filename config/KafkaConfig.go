package config

import (
	"log"

	"github.com/IBM/sarama"
)

func InitKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Println("Error al crear el productor de Kafka:", err)
		return nil, err
	}

	return producer, nil
}

func EnviarMensajeKafka(producer sarama.SyncProducer, mensaje string, topic string) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(mensaje),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Println("Error al enviar mensaje a Kafka:", err)
		return err
	}

	log.Printf("Mensaje enviado a Kafka topic %s [partition %d, offset %d]\n", topic, partition, offset)
	return nil
}
