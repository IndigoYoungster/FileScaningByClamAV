package kafka

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/internal/filesender"
	"github.com/segmentio/kafka-go"
)

func (k *Kafka) Consumer(ctx context.Context, configFolder string) {
	config := kafka.ReaderConfig{
		Brokers:       k.Config.KafkaConsumer.Brokers,
		Topic:         k.Config.KafkaConsumer.Topic,
		Partition:     k.Config.KafkaConsumer.Partition,
		QueueCapacity: k.Config.KafkaConsumer.QueueCapacity,
		MinBytes:      5,
		MaxBytes:      10e6,
		MaxWait:       3 * time.Second,
		GroupID:       k.Config.KafkaConsumer.GroupID,
		StartOffset:   kafka.LastOffset,
	}

	configFileSender := filesender.NewConfig(configFolder)
	fileSender := filesender.NewSender(configFileSender)
	fileSender.СreateFolder(fileSender.Config.UploadFolder)

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		mes, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error while receiving message: %s", err.Error())
			continue
		}

		file, err := os.CreateTemp(fileSender.Config.UploadFolder, "*.zip")
		if err != nil {
			log.Printf("Error os.Create temp: %v\n", err.Error())
		}

		_, err = file.Write(mes.Value)
		if err != nil {
			log.Printf("Error: file.Write: %v\n", err.Error())
		}

		log.Printf("Message at topic: \t%v\nMessage at partition: \t%v\nMessage at offset: \t%v\n", mes.Topic, mes.Partition, mes.Offset)

		file.Close()

		fileSender.Start(file)
	}
}
