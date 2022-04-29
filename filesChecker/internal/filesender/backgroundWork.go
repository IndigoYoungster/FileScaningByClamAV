package filesender

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

const sendFilesCount = 4

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

	for _, file := range files {
		sendFilesToScan(folder, file.Name())
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
