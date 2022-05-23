package apiserver

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	var bodyContent []byte
	r.Body.Read(bodyContent)
	r.Body.Close()
	fmt.Println(bodyContent)
	fmt.Fprint(w, "Ping correctly!")
}
