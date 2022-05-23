package filesender

import (
	"bytes"
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

func (s *sender) sendFileToDb(file *os.File, params *models.Params) {
	var buf bytes.Buffer
	multipartWriter := multipart.NewWriter(&buf)

	file, err := os.Open(file.Name())
	check(err)
	defer os.Remove(file.Name())
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Error fileOpen.Stat()\nMessage: %v\n", err.Error())
	}
	fileWriter, err := multipartWriter.CreateFormFile("file", fileInfo.Name())
	check(err)

	_, err = io.Copy(fileWriter, file)
	check(err)

	// TODO: сделать в конфиге поле для выбора куда вставлять параметры.
	// For add params to form.
	setParamsToFields(multipartWriter, params)

	// For add params to uri.
	//requestParams := setParamsToUri(params)

	multipartWriter.Close()

	//TODO: change to the working address of the database
	req, err := http.NewRequest("POST", s.Config.TestDbRequest.Uri, &buf)
	//req, err := http.NewRequest("POST", s.Config.TestDbRequest.Uri+requestParams, &buf)
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
