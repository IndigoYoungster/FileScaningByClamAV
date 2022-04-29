package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const uploadFolder = "uploadFiles"

func main() {
	addr := flag.String("addr", ":8082", "network port")
	flag.Parse()

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/ping", ping).Methods("GET")
	myRouter.HandleFunc("/api/upload", uploadFiles).Methods("POST")

	createFolder(uploadFolder)

	fmt.Fprintf(log.Writer(), "Start listen on port %s\n", *addr)
	err := http.ListenAndServe(*addr, myRouter)
	log.Fatal(err)
}
