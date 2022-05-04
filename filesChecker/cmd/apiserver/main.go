package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/internal/filesender"
	"github.com/gorilla/mux"
)

const uploadFolder = "uploadFiles"
const tempPrefix = "temp-"
const tickerDuration = time.Second * 5

func main() {
	addr := flag.String("addr", ":8082", "network port")
	flag.Parse()

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/ping", ping).Methods("GET")
	myRouter.HandleFunc("/api/upload", uploadFiles).Methods("POST")

	createFolder(uploadFolder)

	go filesender.Scheduler(tickerDuration, uploadFolder)

	fmt.Fprintf(log.Writer(), "Start listen on port %s\n", *addr)
	err := http.ListenAndServe(*addr, myRouter)
	log.Fatal(err)
}
