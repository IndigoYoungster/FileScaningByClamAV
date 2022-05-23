package apiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	Config Configuration
	Router *mux.Router
}

func NewApi(config *Configuration) *Api {
	return &Api{
		Config: *config,
		Router: mux.NewRouter().StrictSlash(true),
	}
}

func (api *Api) Start() error {
	api.Router.HandleFunc("/api/ping", ping).Methods("GET")

	addr := api.Config.Network.Ip + ":" + api.Config.Network.Port
	fmt.Printf("Start listen on address %s\n", addr)

	return http.ListenAndServe(addr, api.Router)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
