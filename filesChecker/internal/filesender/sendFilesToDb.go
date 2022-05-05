package filesender

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func sendFilesToDb(folder string, fileName string) {
	var buf bytes.Buffer
	multipartWriter := multipart.NewWriter(&buf)

	file, err := os.Open(folder + "/" + fileName)
	check(err)
	defer file.Close()

	fileWriter, err := multipartWriter.CreateFormFile("file", fileName)
	check(err)

	_, err = io.Copy(fileWriter, file)
	check(err)

	multipartWriter.Close()

	//TODO: change to the working address of the database
	req, err := http.NewRequest("POST", "http://localhost:8083/api/test", &buf)
	check(err)

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	timeout := time.Duration(time.Second * 5)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	check(err)
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}
