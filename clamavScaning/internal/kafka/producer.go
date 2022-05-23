package kafka

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func (k *Kafka) producer(ctx context.Context, file *os.File) error {
	// w := &kafka.Writer{
	// 	Addr:     kafka.TCP("127.0.0.1:9092"),
	// 	Topic:    "files-from-scan-topic",
	// 	Balancer: &kafka.LeastBytes{},
	// }

	w := &kafka.Writer{
		Addr:     kafka.TCP(k.Config.KafkaProducer.Addr),
		Topic:    k.Config.KafkaProducer.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	file, err := os.Open(file.Name())
	defer os.Remove(file.Name())
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Printf("Error: file.Stat()\nMessage: %v\n", err.Error())
		return err
	}

	targetFileBytes, err := os.ReadFile(file.Name())
	if err != nil {
		log.Printf("Error: os.ReadFile(targetFile.Name())\nMessage: %v\n", err.Error())
		return err
	}

	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(fileStat.Name()),
		Value: targetFileBytes,
	})

	if err != nil {
		log.Printf("Error: kafka.WriteError\nMessage: %v\n", err.Error())
		return err
	}

	if err := w.Close(); err != nil {
		log.Fatalf("Error: failed to close kafka writer\nMessage: %v\n", err.Error())
		return err
	}

	return nil
}
