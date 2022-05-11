package main

import (
	"flag"
	"log"

	"github.com/IndigoYoungster/FileScaningByClamAV/clamavScaning/internal/apiserver"
)

func main() {
	configFolder := flag.String("conf", "configs/", "configurations folder")
	flag.Parse()

	config := apiserver.NewConfig(*configFolder)
	api := apiserver.NewApi(config)

	err := api.Start()
	log.Fatal(err)
}
