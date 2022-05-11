package apiserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Api struct {
	Config *Configuration
	Router *mux.Router
}

func NewApi(config *Configuration) *Api {
	api := &Api{
		Config: config,
		Router: mux.NewRouter().StrictSlash(true),
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
