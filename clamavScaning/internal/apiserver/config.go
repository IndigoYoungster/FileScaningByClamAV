package apiserver

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Network   network   `yaml:"network"`
	ClamavApi clamavApi `yaml:"clamavApi"`
}

type network struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type clamavApi struct {
	Uri string `yaml:"uri"`
}

func NewConfig(folderPath string) (config *Configuration) {
	fileName := "appSettings.yaml"

	if !strings.HasSuffix(folderPath, "/") {
		folderPath = folderPath + "/"
	}

	filePath := folderPath + fileName

	yfile, err := ioutil.ReadFile(filePath)
	check(err)

	config = new(Configuration)
	err = yaml.Unmarshal(yfile, config)
	check(err)

	return config
}
