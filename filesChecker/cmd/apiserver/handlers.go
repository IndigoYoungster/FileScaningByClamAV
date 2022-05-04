package main

import (
	"fmt"
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ping is correct!")
}

func uploadFiles(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(2048)
	check(err)

	file, fileHeader, err := r.FormFile("file")
	check(err)
	defer file.Close()

	localZipTemp := createZipTemp(file, fileHeader.Filename)
	defer os.Remove(localZipTemp.Name())

	jsonFileTemp := savingParams(r, localZipTemp.Name())
	defer os.Remove(jsonFileTemp.Name())

	saveParamsAndZip(localZipTemp.Name(), jsonFileTemp)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
