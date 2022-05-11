package apiserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	UploadFolder string  `yaml:"uploadFolder"`
	TempPostfix  string  `yaml:"tempPostfix"`
	Network      network `yaml:"network"`
}

type network struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

func (c *Configuration) String() string {
	return fmt.Sprintf("UploadFolder : %s\nTempPostfix : %s\nNetwork : %v\n", c.UploadFolder, c.TempPostfix, c.Network)
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

	log.Println(config.String())
	return config
}
