package filesender

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/models"
)

func (s *sender) sendFilesToScan(folder string, fileNames []string) *models.ResponseModel {
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

		file.Close()
	}
	multipartWriter.Close()

	req, err := http.NewRequest("POST", s.config.ScanFiles.Uri, &b)
	check(err)

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	timeout := time.Duration(time.Second * 10)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	var respModel models.ResponseModel
	err = json.NewDecoder(resp.Body).Decode(&respModel)
	check(err)

	// FOR PRINT RESPONSE MODEL ---->
	log.Print(respModel.String())

	return &respModel
}
