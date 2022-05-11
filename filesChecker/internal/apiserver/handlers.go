package apiserver

import (
	"fmt"
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ping is correct!")
}

func (api *Api) uploadFiles(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	check(err)

	file, fileHeader, err := r.FormFile("file")
	check(err)
	defer file.Close()

	localZipTemp := api.createZipTemp(file, fileHeader.Filename)
	defer os.Remove(localZipTemp.Name())
	defer localZipTemp.Close()

	jsonFileTemp := savingParams(r, localZipTemp.Name())
	defer os.Remove(jsonFileTemp.Name())
	defer jsonFileTemp.Close()

	api.saveParamsAndZip(localZipTemp.Name(), jsonFileTemp)
}
