package apiserver

import (
	"archive/zip"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/models"
)

func (api *Api) createZipTemp(file multipart.File, fileName string) *os.File {
	localZipTemp, err := os.Create(api.Config.TempFolder + "/" + fileName)
	check(err)
	defer localZipTemp.Close()

	_, err = io.Copy(localZipTemp, file)
	check(err)

	return localZipTemp
}

func savingParams(r *http.Request, localZipTempName string) *os.File {
	parameters := new(models.Params)

	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		parameters.RemoteAddr = forwarded
	} else {
		parameters.RemoteAddr = r.RemoteAddr
	}
	parameters.App = r.FormValue("app")
	parameters.Usr = r.FormValue("usr")
	parameters.Comp = r.FormValue("comp")

	paramsFileName := strings.Replace(localZipTempName, ".zip", ".json", 1)
	jsonFile, err := os.Create(paramsFileName)

	check(err)
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(parameters)
	check(err)

	return jsonFile
}

func (api *Api) saveParamsAndZip(localZipTempName string, file *os.File) *os.File {
	jsonFile, err := os.Open(file.Name())
	check(err)
	defer jsonFile.Close()

	fileInfo, err := jsonFile.Stat()
	check(err)

	zipReader, err := zip.OpenReader(localZipTempName)
	check(err)
	defer zipReader.Close()

	targetFile, err := os.Create(strings.Replace(localZipTempName, ".zip", api.Config.TempPostfix+".zip", 1))
	check(err)
	defer targetFile.Close()

	targetZipWriter := zip.NewWriter(targetFile)
	defer targetZipWriter.Close()

	for _, zipItem := range zipReader.File {
		targetZipWriter.Copy(zipItem)
	}

	fileBytes, err := os.ReadFile(jsonFile.Name())
	check(err)
	targetItem, err := targetZipWriter.Create(fileInfo.Name())
	check(err)
	_, err = targetItem.Write(fileBytes)
	check(err)

	err = targetZipWriter.Flush()
	check(err)

	return targetFile
}
