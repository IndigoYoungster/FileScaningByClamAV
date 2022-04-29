package main

import (
	"fmt"
	"io"
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

	localFile, err := os.Create(uploadFolder + "/" + fileHeader.Filename)
	check(err)
	defer localFile.Close()

	io.Copy(localFile, file)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
