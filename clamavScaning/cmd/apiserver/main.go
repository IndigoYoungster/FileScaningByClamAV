package main

import (
	"context"
	"flag"
	"log"

	"github.com/IndigoYoungster/FileScaningByClamAV/clamavScaning/internal/apiserver"
	"github.com/IndigoYoungster/FileScaningByClamAV/clamavScaning/internal/kafka"
)

func main() {
	configFolder := flag.String("conf", "configs/", "configurations folder")
	flag.Parse()

	config := apiserver.NewConfig(*configFolder)
	api := apiserver.NewApi(config)

	configKafka := kafka.NewConfig(*configFolder)
	kaf := kafka.NewKafka(configKafka)

	ctx := context.Background()
	go kaf.Consumer(ctx)

	err := api.Start()
	log.Fatal(err)
}
