package kafka

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	KafkaConsumer kafkaConsumer `yaml:"kafkaConsumer"`
	KafkaProducer kafkaProducer `yaml:"kafkaProducer"`
}

type kafkaConsumer struct {
	Topic         string   `yaml:"topic"`
	Partition     int      `yaml:"partition"`
	QueueCapacity int      `yaml:"queueCapacity"`
	GroupID       string   `yaml:"groupId"`
	Brokers       []string `yaml:"brokers"`
}

type kafkaProducer struct {
	Addr  string `yaml:"addr"`
	Topic string `yaml:"topic"`
}

func NewConfig(folderName string) *Configuration {
	fileName := "kafkaSettings.yaml"

	if !strings.HasSuffix(folderName, "/") {
		folderName += "/"
	}

	filePath := folderName + fileName

	yfile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error with read new kafka config\nMessage: %v\n", err.Error())
	}

	config := &Configuration{}
	err = yaml.Unmarshal(yfile, config)
	if err != nil {
		log.Printf("Error with unmarshal kafka config\nMessage: %v\n", err.Error())
	}
	return config
}

type Kafka struct {
	Config *Configuration
}

func NewKafka(config *Configuration) *Kafka {
	return &Kafka{
		Config: config,
	}
}
