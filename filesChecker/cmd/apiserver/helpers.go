package main

import (
	"archive/zip"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func createFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, 0666)
	}
}

func createZipTemp(file multipart.File, fileName string) *os.File {
	localZipTemp, err := os.Create(uploadFolder + "/" + tempPrefix + fileName)
	check(err)
	defer localZipTemp.Close()

	_, err = io.Copy(localZipTemp, file)
	check(err)

	return localZipTemp
}

func savingParams(r *http.Request, localZipTempName string) *os.File {
	parameters := new(params)

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
	paramsFileName = strings.Replace(paramsFileName, tempPrefix, "", 1)
	jsonFile, err := os.Create(paramsFileName)

	check(err)
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(parameters)
	check(err)

	return jsonFile
}

func saveParamsAndZip(localZipTempName string, file *os.File) {
	jsonFile, err := os.Open(file.Name())
	check(err)
	defer jsonFile.Close()

	fileInfo, err := jsonFile.Stat()
	check(err)

	zipReader, err := zip.OpenReader(localZipTempName)
	check(err)
	defer zipReader.Close()

	targetFile, err := os.Create(strings.Replace(localZipTempName, tempPrefix, "", 1))
	check(err)
	defer targetFile.Close()

	targetZipWriter := zip.NewWriter(targetFile)
	defer targetZipWriter.Close()

	for _, zipItem := range zipReader.File {
		zipItemReader, err := zipItem.Open()
		defer zipItemReader.Close()
		check(err)
		header, err := zip.FileInfoHeader(zipItem.FileInfo())
		check(err)
		header.Name = zipItem.Name
		header.Method = zipItem.Method
		targetItem, err := targetZipWriter.CreateHeader(header)
		check(err)
		_, err = io.Copy(targetItem, zipItemReader)
		check(err)
	}

	fileBytes, err := os.ReadFile(jsonFile.Name())
	check(err)
	targetItem, err := targetZipWriter.Create(fileInfo.Name())
	check(err)
	_, err = targetItem.Write(fileBytes)
	check(err)

	err = targetZipWriter.Flush()
	check(err)
}
