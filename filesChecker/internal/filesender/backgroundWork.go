package filesender

import (
	"log"
	"os"
	"time"
)

const maxSendFilesCount = 4
const tempPrefix = "-temp"

// for test
const sendToDb = false

func Scheduler(d time.Duration, folder string) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		<-ticker.C

		filesToScan := checkNewFilesInFolder(folder)

		if filesToScan != nil {
			responseModel := sendFilesToScan(folder, filesToScan)

			correctFileNames := checkCorrectResponse(folder, responseModel)
			if len(correctFileNames) != 0 {
				for _, fileName := range correctFileNames {
					if sendToDb {
						sendingFileName, params := getFileAndParams(folder, fileName)
						sendFilesToDb(sendingFileName, params)
					} else {
						sendingFileName, params := getFileAndParams(folder, fileName)
						os.Remove(sendingFileName)
						log.Printf("File %s - Correct after check\nParams--->\n%s\n", fileName, params.String())
					}
				}
			}
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
