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
	err := r.ParseMultipartForm(32 << 20)
	check(err)

	file, fileHeader, err := r.FormFile("file")
	check(err)
	defer file.Close()

	localZipTemp := createZipTemp(file, fileHeader.Filename)
	defer os.Remove(localZipTemp.Name())
	defer localZipTemp.Close()

	jsonFileTemp := savingParams(r, localZipTemp.Name())
	defer os.Remove(jsonFileTemp.Name())
	defer jsonFileTemp.Close()

	saveParamsAndZip(localZipTemp.Name(), jsonFileTemp)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
