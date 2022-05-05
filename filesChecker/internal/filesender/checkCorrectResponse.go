package filesender

import (
	"log"
	"os"
)

func checkCorrectResponse(folder string, responseModel *responseModel) (correctFileNames []string) {
	if responseModel.Success {
		for _, result := range responseModel.Data.Result {
			if !result.IsInfected {
				if _, err := os.Stat(folder + "/" + result.Name); err == nil {
					correctFileNames = append(correctFileNames, result.Name)
				} else {
					log.Printf("WARNING: File %s doesn't exist in directory!!!", result.Name)
				}
			} else {
				log.Fatalf("File %s IS INFECTED!!!\n", result.Name)
			}
		}
	} else {
		log.Println("WARNING: Success == false")
	}

	return correctFileNames
}
