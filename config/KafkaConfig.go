package config

import (
	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	BrokerAddresses []string
	SaramaConfig    *sarama.Config
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		BrokerAddresses: []string{"172.0.0.1:9092"}, // Aquí puedes poner las direcciones de tus brokers
		SaramaConfig:    sarama.NewConfig(),         // Configura aquí tu configuración Sarama
	}
}
