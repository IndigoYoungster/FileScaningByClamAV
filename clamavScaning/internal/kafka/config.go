package kafka

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	TempFolder    string        `yaml:"tempFolder"`
	KafkaConsumer kafkaConsumer `yaml:"kafkaConsumer"`
	KafkaProducer kafkaProducer `yaml:"kafkaProducer"`
	ClamavApi     clamavApi     `yaml:"clamavApi"`
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

type clamavApi struct {
	Uri string `yaml:"uri"`
}

func NewConfig(folderPath string) *Configuration {
	fileName := "fileScaningSettings.yaml"

	if !strings.HasSuffix(folderPath, "/") {
		folderPath += "/"
	}

	filePath := folderPath + fileName

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

func createFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, 0666)
	}
}
