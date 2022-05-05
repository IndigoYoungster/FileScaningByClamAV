package filesender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func sendFilesToScan(folder string, fileNames []string) {
	var b bytes.Buffer
	multipartWriter := multipart.NewWriter(&b)

	for _, fileName := range fileNames {
		file, err := os.Open(folder + "/" + fileName)
		check(err)

		fileContents, err := ioutil.ReadAll(file)
		check(err)

		fileWriter, err := multipartWriter.CreateFormFile("file", fileName)
		check(err)

		_, err = fileWriter.Write(fileContents)
		check(err)
	}
	multipartWriter.Close()

	req, err := http.NewRequest("POST", "http://localhost:8081/api/scan", &b)
	check(err)

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	timeout := time.Duration(time.Second * 5)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	var respModel responseModel
	err = json.NewDecoder(resp.Body).Decode(&respModel)
	if err == io.EOF {
		log.Println("JSON decode complete success")
	}
	fmt.Print(respModel.String())
}
