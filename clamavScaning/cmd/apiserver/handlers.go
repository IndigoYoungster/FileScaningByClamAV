package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	var bodyContent []byte
	r.Body.Read(bodyContent)
	r.Body.Close()
	fmt.Println(bodyContent)
	fmt.Fprint(w, "Ping correctly!")
}

func sendToScan(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(2048)
	check(err)

	file, fileHeader, err := r.FormFile("file")
	check(err)
	defer file.Close()

	var requestBody bytes.Buffer

	multipartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multipartWriter.CreateFormFile("FILES", fileHeader.Filename)
	check(err)

	_, err = io.Copy(fileWriter, file)
	check(err)

	multipartWriter.Close()

	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/scan", &requestBody)
	check(err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)
	check(err)
	defer response.Body.Close()

	//FOR SEND JSON ---->
	io.Copy(w, response.Body)

	//FOR SEND STRING ---->
	//var respModel responseModel
	//err = json.NewDecoder(response.Body).Decode(&respModel)
	//check(err)
	//fmt.Fprintln(w, respModel.String())

	//FOR PRINT RESPONSE ---->
	//io.Copy(os.Stdout, response.Body)

	//FOR SAVING FILES ---->
	//saveFiles(fileHeader.Filename, file)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}

func saveFiles(filename string, file multipart.File) {
	fileBytes, err := ioutil.ReadAll(file)
	check(err)
	err = ioutil.WriteFile(filename, fileBytes, 0666)
	check(err)
}
