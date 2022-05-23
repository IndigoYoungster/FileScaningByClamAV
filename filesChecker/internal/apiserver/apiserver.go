package apiserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/internal/kafka"
	"github.com/gorilla/mux"
)

type Api struct {
	Config     *Configuration
	Router     *mux.Router
	Kafka      *kafka.Kafka
	FileChanel chan *os.File
}

func NewApi(config *Configuration, kafConfig *kafka.Configuration) *Api {
	api := &Api{
		Config:     config,
		Router:     mux.NewRouter().StrictSlash(true),
		Kafka:      kafka.NewKafka(kafConfig),
		FileChanel: make(chan *os.File, 3),
	}
	return api
}

func (api *Api) Start() error {
	api.Router.HandleFunc("/api/ping", ping).Methods("GET")
	api.Router.HandleFunc("/api/upload", api.uploadFiles).Methods("POST")

	addr := api.Config.Network.Ip + ":" + api.Config.Network.Port
	fmt.Printf("Start listen on port %s\n", addr)

	return http.ListenAndServe(addr, api.Router)
}

func CreateFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, 0666)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
