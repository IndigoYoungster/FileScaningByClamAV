package filesender

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func sendFilesToScan(folder string, fileName string) {
	file, err := os.Open(folder + "/" + fileName)
	check(err)

	fileContents, err := ioutil.ReadAll(file)
	check(err)

	var b bytes.Buffer
	multipartWriter := multipart.NewWriter(&b)

	fileWriter, err := multipartWriter.CreateFormFile("file", fileName)
	check(err)

	_, err = fileWriter.Write(fileContents)
	check(err)
	multipartWriter.Close()

	resp, err := http.Post("http://localhost:8081/api/scan", multipartWriter.FormDataContentType(), &b)
	check(err)
	defer resp.Body.Close()

	var respModel responseModel
	err = json.NewDecoder(resp.Body).Decode(&respModel)
	check(err)

	//FOR PRINT RESULT ---->
	io.Copy(os.Stdout, &respModel)
}
