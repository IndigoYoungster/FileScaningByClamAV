package filesender

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	UploadFolder  string        `yaml:"uploadFolder"`
	CrashDb       crashDb       `yaml:"crashDb"`
	TestDbRequest testDbRequest `yaml:"testDbRequest"`
}

type scanFiles struct {
	Uri string `yaml:"uri"`
}

type crashDb struct {
	Uri string `yaml:"uri"`
}

type testDbRequest struct {
	Uri string `yaml:"uri"`
}

func (c *Configuration) String() string {
	return fmt.Sprintf("UploadFolder: %s\nCrashDb: %v\nTestDbRequest: %v\n", c.UploadFolder, c.CrashDb, c.TestDbRequest)
}

func NewConfig(folderPath string) (config *Configuration) {
	fileName := "backgroundSettings.yaml"

	if !strings.HasSuffix(folderPath, "/") {
		folderPath = folderPath + "/"
	}

	filePath := folderPath + fileName

	yfile, err := ioutil.ReadFile(filePath)
	check(err)

	config = new(Configuration)
	err = yaml.Unmarshal(yfile, config)
	check(err)

	log.Println(config.String())
	return config
}
