package filesender

import (
	"archive/zip"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/IndigoYoungster/FileScaningByClamAV/filesChecker/models"
)

func (s *sender) reExtractFile(file *os.File) (*os.File, *models.Params) {
	zipFile, err := zip.OpenReader(file.Name())
	if err != nil {
		log.Printf("Error zip.OpenReader\nMessage: %v\n", err.Error())
	}
	defer os.Remove(file.Name())
	defer zipFile.Close()

	var jsonFileName string
	params := &models.Params{}
	jsonFileName, params = s.getFileNameAndParams(zipFile)

	targetZip, err := os.Create(strings.Replace(jsonFileName, ".json", ".zip", 1))
	if err != nil {
		log.Printf("Error os.Create()\nMessage: %v\n", err.Error())
	}
	defer targetZip.Close()

	targetZipWriter := zip.NewWriter(targetZip)
	defer targetZipWriter.Close()

	for _, zipItem := range zipFile.File {
		if strings.HasSuffix(zipItem.Name, ".json") {
			continue
		}

		err = targetZipWriter.Copy(zipItem)
		if err != nil {
			log.Printf("Error targetZipWriter.Copy(zipItem)\nMessage: %v\n", err.Error())
		}
	}

	return targetZip, params
}

func (s *sender) getFileNameAndParams(zipFile *zip.ReadCloser) (string, *models.Params) {
	var jsonFileName string
	params := &models.Params{}
	for _, zipItem := range zipFile.File {
		fileName := zipItem.Name
		if strings.HasSuffix(fileName, ".json") {
			jsonFileName = fileName

			zipItemRC, err := zipItem.Open()
			if err != nil {
				log.Printf("Error zipItem.Open()\nMessage: %v\n", err.Error())
			}

			err = json.NewDecoder(zipItemRC).Decode(params)
			if err != nil {
				log.Printf("Error json.NewDecoder()\nMessage: %v\n", err.Error())
			}

			zipItemRC.Close()
			break
		}
	}

	if jsonFileName == "" {
		log.Println("Warning: json does not exist. [getFilenameAndParams]")
	}
	if params == nil {
		log.Println("Warning: params not created. [getFilenameAndParams]")
	}

	tempDir := s.Config.UploadFolder
	return tempDir + "/" + jsonFileName, params
}
