package filesender

import (
	"fmt"
	"io/ioutil"
)

func checkNewFilesInFolder(folder string) []string {
	files, err := ioutil.ReadDir(folder)
	check(err)

	if len(files) == 0 {
		fmt.Println("folder empty")
		return nil
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
	return filesToScan
}
