package filesender

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	MaxSendFilesCount int           `yaml:"maxSendFilesCount"`
	TickerDuration    int           `yaml:"tickerDuration"`
	TempPostfix       string        `yaml:"tempPostfix"`
	ScanFiles         scanFiles     `yaml:"scanFiles"`
	CrashDb           crashDb       `yaml:"crashDb"`
	TestDbRequest     testDbRequest `yaml:"testDbRequest"`
}

type scanFiles struct {
	Uri string `yaml:"uri"`
}

type crashDb struct {
	SendToDb bool   `yaml:"sendToDb"`
	Uri      string `yaml:"uri"`
}

type testDbRequest struct {
	Uri string `yaml:"uri"`
}

func (c *Configuration) String() string {
	return fmt.Sprintf("MaxSendFilesCount: %d\nTickerDuration: %d\nTempPostfix: %s\nScanFiles: %v\nCrashDb: %v\nTestDbRequest: %v\n", c.MaxSendFilesCount, c.TickerDuration, c.TempPostfix, c.ScanFiles, c.CrashDb, c.TestDbRequest)
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
