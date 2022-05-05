package filesender

import (
	"log"
	"time"
)

const maxSendFilesCount = 4

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
					sendFilesToDb(folder, fileName)
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
