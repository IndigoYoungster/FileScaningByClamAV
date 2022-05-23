package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/internal/apiserver"
	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/internal/kafka"
)

func main() {
	configFolder := flag.String("conf", "configs/", "configurations folder")
	flag.Parse()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()

	configKafka := kafka.NewConfig(*configFolder)
	kaf := kafka.NewKafka(configKafka)
	go kaf.Consumer(ctx, *configFolder)

	config := apiserver.NewConfig(*configFolder)
	api := apiserver.NewApi(config, configKafka)
	go kaf.Producer(ctx, api.FileChanel)

	apiserver.CreateFolder(api.Config.TempFolder)

	err := api.Start()
	log.Fatal(err)
}
