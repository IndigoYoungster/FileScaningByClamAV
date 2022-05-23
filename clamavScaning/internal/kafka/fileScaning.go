package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/IndigoYoungster/FileScaningByClamAV/clamavScaning/models"
)

func (k *Kafka) sendToScan(ctx context.Context, file *os.File) {
	respModel := k.send(file)

	success := checkResponse(respModel)
	if !success {
		log.Printf("File: %s\n", file.Name())
		os.Remove(file.Name())
		return
	}

	if respModel.Data.Result[0].IsInfected {
		log.Printf("WARNING: File - %s IS INFECTED\nDeleting......\n", file.Name())
		os.Remove(file.Name())
		return
	}

	log.Println("SCAN SUCCESS!")
	err := k.producer(ctx, file)
	if err != nil {
		log.Printf("Error with file producer\nFile - %s\n", file.Name())
	} else {
		log.Printf("Scaned file sending success. File - %s\n", file.Name())
	}
}

func checkResponse(respModel *models.ResponseModel) bool {
	if !respModel.Success {
		log.Println("Error: response from clamd - false")
		return false
	}
	if respModel.Data.Result == nil {
		log.Println("Error: result is nil")
		return false
	}

	return true
}

func (k *Kafka) send(file *os.File) *models.ResponseModel {
	var b bytes.Buffer
	multipartWriter := multipart.NewWriter(&b)

	fileOpen, err := os.Open(file.Name())
	if err != nil {
		log.Printf("Error os.Open(file.Name())\nMessage: %v\n", err.Error())
	}

	fileStat, err := fileOpen.Stat()
	if err != nil {
		log.Printf("Error fileOpen.Stat()\nMessage: %v\n", err.Error())
	}

	fileWriter, err := multipartWriter.CreateFormFile("FILES", fileStat.Name())
	if err != nil {
		log.Printf("Error multipartWriter.CreateFormFile\nMessage: %v\n", err.Error())
	}

	_, err = io.Copy(fileWriter, fileOpen)
	if err != nil {
		log.Printf("Error io.Copy(fileWriter, file)\nMessage: %v\n", err.Error())
	}

	fileOpen.Close()
	multipartWriter.Close()

	req, err := http.NewRequest("POST", k.Config.ClamavApi.Uri, &b)
	if err != nil {
		log.Printf("Error http.NewRequest(Clamav.Uri)\nMessage: %v\n", err.Error())
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	timeout := time.Duration(time.Second * 10)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error client.Do(request)\nMessage: %v\n", err.Error())
	}
	defer resp.Body.Close()

	respModel := &models.ResponseModel{}
	err = json.NewDecoder(resp.Body).Decode(respModel)
	if err != nil {
		log.Printf("Error json.NewDecoder(resp.Body)\nMessage: %v\n", err.Error())
	}
	return respModel
}
