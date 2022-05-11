package filesender

import (
	"log"
	"os"
)

func (s *sender) Schedule(folder string) {
	filesToScan := s.checkNewFilesInFolder(folder)

	if filesToScan != nil {
		responseModel := s.sendFilesToScan(folder, filesToScan)

		correctFileNames := checkCorrectResponse(folder, responseModel)
		if len(correctFileNames) != 0 {
			for _, fileName := range correctFileNames {
				if s.config.CrashDb.SendToDb {
					sendingFileName, params := s.getFileAndParams(folder, fileName)
					s.sendFilesToDb(sendingFileName, params)
				} else {
					sendingFileName, params := s.getFileAndParams(folder, fileName)
					os.Remove(sendingFileName)
					log.Printf("File %s - Correct after check\nParams--->\n%s\n", fileName, params.String())
				}
			}
		}
	}
}
