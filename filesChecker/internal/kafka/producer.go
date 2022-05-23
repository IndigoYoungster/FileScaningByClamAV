package kafka

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func (k *Kafka) Producer(ctx context.Context, fileChanel chan *os.File) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.Config.KafkaProducer.Addr),
		Topic:    k.Config.KafkaProducer.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()

	for {
		file := <-fileChanel
		file, err := os.Open(file.Name())
		if err != nil {
			log.Printf("Error: producer os.Open()\nMessage: %v\n", err.Error())
		}

		fileStat, err := file.Stat()
		if err != nil {
			log.Printf("Error: file.Stat()\nMessage: %v\n", err.Error())
		}

		targetFileBytes, err := os.ReadFile(file.Name())
		if err != nil {
			log.Printf("Error: os.ReadFile(targetFile.Name())\nMessage: %v\n", err.Error())
		}

		err = w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(fileStat.Name()),
			Value: targetFileBytes,
		})

		if err != nil {
			log.Printf("Error: kafka.WriteError\nMessage: %v\n", err.Error())
		}

		log.Printf("SUCCESS send to scan - %s\n", fileStat.Name())

		if err := file.Close(); err != nil {
			log.Printf("Error: producer file.Close()\nMessage: %v\n", err.Error())
		}
		if err := os.Remove(file.Name()); err != nil {
			log.Printf("Error: producer os.Remove\nMessage: %v\n", err.Error())
		}
	}
}
