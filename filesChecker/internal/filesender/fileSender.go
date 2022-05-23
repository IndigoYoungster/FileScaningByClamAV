package filesender

import (
	"log"
	"os"
)

type sender struct {
	Config *Configuration
}

func NewSender(config *Configuration) *sender {
	return &sender{
		Config: config,
	}
}

func (s *sender) Start(file *os.File) {
	reExtractFile, params := s.reExtractFile(file)

	log.Printf("SUCCESS reExtract file - %s\nParams: %s\n", reExtractFile.Name(), params.String())
	s.sendFileToDb(reExtractFile, params)
}

func (s *sender) СreateFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, 0666)
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
