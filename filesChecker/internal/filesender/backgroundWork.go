package filesender

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const maxSendFilesCount = 4

func Scheduler(d time.Duration, folder string) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		<-ticker.C

		checkNewFilesInFolder(folder)
	}
}

func checkNewFilesInFolder(folder string) {
	files, err := ioutil.ReadDir(folder)
	check(err)

	if len(files) == 0 {
		fmt.Println("folder empty")
		return
	}

	count := 0
	var filesToScan []string
	for _, file := range files {
		filesToScan = append(filesToScan, file.Name())

		count++
		if count >= maxSendFilesCount {
			break
		}
	}
	sendFilesToScan(folder, filesToScan)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
