package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	addr := flag.String("addr", ":8081", "network port")
	flag.Parse()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/ping", ping).Methods("GET")
	myRouter.HandleFunc("/api/scan", sendToScan).Methods("POST")

	fmt.Fprintf(os.Stdout, "Start listen on port %s\n", *addr)
	err := http.ListenAndServe(*addr, myRouter)
	log.Fatal(err)
}
