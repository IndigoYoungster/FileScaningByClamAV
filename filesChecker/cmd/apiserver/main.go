package main

import (
	"flag"
	"log"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/internal/apiserver"
	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/internal/filesender"
)

func main() {
	configFolder := flag.String("conf", "configs/", "configurations folder")
	flag.Parse()

	config := apiserver.NewConfig(*configFolder)
	api := apiserver.NewApi(config)

	apiserver.CreateFolder(api.Config.UploadFolder)

	configFileSender := filesender.NewConfig(*configFolder)
	fileSender := filesender.NewSender(configFileSender)

	go fileSender.Start(api.Config.UploadFolder)

	err := api.Start()
	log.Fatal(err)
}
