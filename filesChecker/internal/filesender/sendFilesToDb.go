package filesender

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/models"
)

func (s *sender) sendFilesToDb(fileName string, params *models.Params) {
	var buf bytes.Buffer
	multipartWriter := multipart.NewWriter(&buf)

	file, err := os.Open(fileName)
	check(err)
	defer os.Remove(file.Name())
	defer file.Close()

	fileWriter, err := multipartWriter.CreateFormFile("file", fileName)
	check(err)

	_, err = io.Copy(fileWriter, file)
	check(err)

	// for add params to form
	setParamsToFields(multipartWriter, params)

	// for add params to uri
	//requestParams := setParamsToUri(params)

	multipartWriter.Close()

	//TODO: change to the working address of the database
	req, err := http.NewRequest("POST", s.config.CrashDb.Uri, &buf)
	//req, err := http.NewRequest("POST", s.config.CrashDb.Uri+requestParams, &buf)
	check(err)

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	if params.RemoteAddr != "" {
		req.Header.Set("X-FORWARDED-FOR", params.RemoteAddr)
	}

	timeout := time.Duration(time.Second * 10)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	check(err)
	bodyString := string(bodyBytes)
	log.Printf("Status code : %s\nBody: %s\n", resp.Status, bodyString)
}

func (s *sender) getFileAndParams(folder, fileName string) (zipFileName string, params *models.Params) {
	zipReader, err := zip.OpenReader(folder + "/" + fileName)
	check(err)
	defer os.Remove(folder + "/" + fileName)
	defer zipReader.Close()

	targetZipFile, err := os.Create(folder + "/" + strings.Replace(fileName, s.config.TempPostfix, "", 1))
	check(err)
	defer targetZipFile.Close()

	targetZipWriter := zip.NewWriter(targetZipFile)
	defer targetZipWriter.Close()

	targetJsonName := strings.Replace(fileName, s.config.TempPostfix+".zip", ".json", 1)
	params = new(models.Params)

	for _, zipItem := range zipReader.File {
		if zipItem.Name == targetJsonName {
			fileRCJson, err := zipItem.Open()
			check(err)

			err = json.NewDecoder(fileRCJson).Decode(params)
			check(err)

			fileRCJson.Close()
			continue
		}
		err = targetZipWriter.Copy(zipItem)
		check(err)
	}

	return targetZipFile.Name(), params
}

func setParamsToFields(multipartWriter *multipart.Writer, params *models.Params) {
	if params.App != "" {
		fieldWriter, err := multipartWriter.CreateFormField("app")
		check(err)

		_, err = fieldWriter.Write([]byte(params.App))
		check(err)
	}
	if params.Comp != "" {
		fieldWriter, err := multipartWriter.CreateFormField("comp")
		check(err)

		_, err = fieldWriter.Write([]byte(params.App))
		check(err)
	}
	if params.Usr != "" {
		fieldWriter, err := multipartWriter.CreateFormField("usr")
		check(err)

		_, err = fieldWriter.Write([]byte(params.App))
		check(err)
	}
}

func setParamsToUri(params *models.Params) string {
	requestParams := make([]string, 3)

	if params.App != "" {
		requestParams = append(requestParams, fmt.Sprintf("app=%s", params.App))
	}
	if params.Comp != "" {
		requestParams = append(requestParams, fmt.Sprintf("comp=%s", params.Comp))
	}
	if params.Usr != "" {
		requestParams = append(requestParams, fmt.Sprintf("usr=%s", params.Usr))
	}

	return "?" + strings.Join(requestParams, "&")
}
