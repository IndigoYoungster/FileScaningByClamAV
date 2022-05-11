package apiserver

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

func ping(w http.ResponseWriter, r *http.Request) {
	var bodyContent []byte
	r.Body.Read(bodyContent)
	r.Body.Close()
	fmt.Println(bodyContent)
	fmt.Fprint(w, "Ping correctly!")
}

func (api *Api) sendToScan(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) //32 MB
	check(err)

	var b bytes.Buffer
	multipartWriter := multipart.NewWriter(&b)

	fileHeaders := r.MultipartForm.File["file"]
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		check(err)

		fileWriter, err := multipartWriter.CreateFormFile("FILES", fileHeader.Filename)
		check(err)

		_, err = io.Copy(fileWriter, file)
		check(err)

		file.Close()
	}
	multipartWriter.Close()

	req, err := http.NewRequest("POST", api.Config.ClamavApi.Uri, &b)
	check(err)

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	timeout := time.Duration(time.Second * 10)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	io.Copy(w, resp.Body)

	// FOR PRINT JSON ---->
	// var respModel responseModel
	// err = json.NewDecoder(resp.Body).Decode(&respModel)
	// check(err)
	// fmt.Println(respModel.String())
}
